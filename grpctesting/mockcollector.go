/*
 * Copyright (c) The go-kit Authors
 */

package grpctesting

import (
	"context"
	"fmt"
	"net"
	"runtime"
	"strings"
	"sync"
	"time"

	collectormetricpb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	metricpb "go.opentelemetry.io/proto/otlp/metrics/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// MakeMockCollector it has been lifted from the https://github.com/open-telemetry/opentelemetry-go with some little tweak
// This will be useful until opentelemetry go release a metrics test library
// TODO delete this file when opentelemetry-go release the metrics library and test framework
func MakeMockCollector(mockConfig *MockConfig) *MockCollector {
	return &MockCollector{
		metricSvc: &mockMetricService{
			storage: NewMetricsStorage(),
			errors:  mockConfig.Errors,
		},
	}
}

type mockMetricService struct {
	collectormetricpb.UnimplementedMetricsServiceServer

	requests int
	errors   []error

	headers metadata.MD
	mu      sync.RWMutex
	storage MetricsStorage
	delay   time.Duration
}

func (mms *mockMetricService) getHeaders() metadata.MD {
	mms.mu.RLock()
	defer mms.mu.RUnlock()
	return mms.headers
}

func (mms *mockMetricService) getMetrics() []*metricpb.Metric {
	mms.mu.RLock()
	defer mms.mu.RUnlock()
	return mms.storage.GetMetrics()
}

func (mms *mockMetricService) Export(ctx context.Context, exp *collectormetricpb.ExportMetricsServiceRequest) (*collectormetricpb.ExportMetricsServiceResponse, error) {
	if mms.delay > 0 {
		time.Sleep(mms.delay)
	}

	mms.mu.Lock()
	defer func() {
		mms.requests++
		mms.mu.Unlock()
	}()

	reply := &collectormetricpb.ExportMetricsServiceResponse{}
	if mms.requests < len(mms.errors) {
		idx := mms.requests
		return reply, mms.errors[idx]
	}

	mms.headers, _ = metadata.FromIncomingContext(ctx)
	mms.storage.AddMetrics(exp)
	return reply, nil
}

// MockCollector is an opentelemetry collector suitable for tests
type MockCollector struct {
	metricSvc *mockMetricService
	endpoint  string
	ln        *listener
	stopFunc  func()
	stopOnce  sync.Once
}

// GetMetrics returns the list of metrics
func (mc *MockCollector) GetMetrics() []*metricpb.Metric {
	return mc.getMetrics()
}

// Stop the collector
func (mc *MockCollector) Stop() error {
	return mc.stop()
}

// GetEndPoint returns the collector Endpoint
func (mc *MockCollector) GetEndPoint() string {
	return mc.endpoint
}

// MockConfig is the collector configuration
type MockConfig struct {
	Errors   []error
	Endpoint string
}

var _ collectormetricpb.MetricsServiceServer = (*mockMetricService)(nil)

var errAlreadyStopped = fmt.Errorf("already stopped")

func (mc *MockCollector) stop() error {
	var err = errAlreadyStopped
	mc.stopOnce.Do(func() {
		err = nil
		if mc.stopFunc != nil {
			mc.stopFunc()
		}
	})
	// Give it sometime to shutdown.
	<-time.After(160 * time.Millisecond)

	// Wait for services to finish reading/writing.
	// Getting the lock ensures the metricSvc is done flushing.
	mc.metricSvc.mu.Lock()
	defer mc.metricSvc.mu.Unlock()
	return err
}

func (mc *MockCollector) GetHeaders() metadata.MD {
	return mc.metricSvc.getHeaders()
}

func (mc *MockCollector) getMetrics() []*metricpb.Metric {
	return mc.metricSvc.getMetrics()
}

// RunMockCollector is a helper function to create a mock Collector
func RunMockCollector() (*MockCollector, error) {
	return RunMockCollectorAtEndpoint("localhost:0")
}

// RunMockCollectorAtEndpoint creates an instance of the MockCollector and starts it
// at the given Endpoint
func RunMockCollectorAtEndpoint(endpoint string) (*MockCollector, error) {
	return RunMockCollectorWithConfig(&MockConfig{Endpoint: endpoint})
}

// RunMockCollectorWithConfig creates an instance of the MockCollector and starts it given
// a mock config
func RunMockCollectorWithConfig(mockConfig *MockConfig) (*MockCollector, error) {
	ln, err := net.Listen("tcp", mockConfig.Endpoint)
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()
	mc := MakeMockCollector(mockConfig)
	collectormetricpb.RegisterMetricsServiceServer(srv, mc.metricSvc)
	mc.ln = newListener(ln)
	go func() {
		_ = srv.Serve((net.Listener)(mc.ln))
	}()

	mc.endpoint = ln.Addr().String()
	// srv.Stop calls Close on mc.ln.
	mc.stopFunc = srv.Stop

	return mc, nil
}

type listener struct {
	closeOnce sync.Once
	wrapped   net.Listener
	C         chan struct{}
}

func newListener(wrapped net.Listener) *listener {
	return &listener{
		wrapped: wrapped,
		C:       make(chan struct{}, 1),
	}
}

func (l *listener) Close() error { return l.wrapped.Close() }

func (l *listener) Addr() net.Addr { return l.wrapped.Addr() }

// Accept waits for and returns the next connection to the listener. It will
// send a signal on l.C that a connection has been made before returning.
func (l *listener) Accept() (net.Conn, error) {
	conn, err := l.wrapped.Accept()
	if err != nil {
		// Go 1.16 exported net.ErrClosed that could clean up this check, but to
		// remain backwards compatible with previous versions of Go that we
		// support the following string evaluation is used instead to keep in line
		// with the previously recommended way to check this:
		// https://github.com/golang/go/issues/4373#issuecomment-353076799
		if strings.Contains(err.Error(), "use of closed network connection") {
			// If the listener has been closed, do not allow callers of
			// WaitForConn to wait for a connection that will never come.
			l.closeOnce.Do(func() { close(l.C) })
		}
		return conn, err
	}

	select {
	case l.C <- struct{}{}:
	default:
		// If C is full, assume nobody is listening and move on.
	}
	return conn, nil
}

// WaitForConn will wait indefintely for a connection to be estabilished with
// the listener before returning.
func (l *listener) WaitForConn() {
	for {
		select {
		case <-l.C:
			return
		default:
			runtime.Gosched()
		}
	}
}

// Collector is an interface that mock collectors should implement,
// so they can be used for the end-to-end testing.
// The code has been lifted from the https://github.com/open-telemetry/opentelemetry-go
// because we cannot import an internal package in go
type Collector interface {
	Stop() error
	GetMetrics() []*metricpb.Metric
	GetHeaders() metadata.MD
	GetEndPoint() string
}

// MetricsStorage stores the metrics. Mock collectors could use it to
// store metrics they have received.
type MetricsStorage struct {
	metrics []*metricpb.Metric
}

// NewMetricsStorage creates a new metrics storage.
func NewMetricsStorage() MetricsStorage {
	return MetricsStorage{}
}

// AddMetrics adds metrics to the metrics storage.
func (s *MetricsStorage) AddMetrics(request *collectormetricpb.ExportMetricsServiceRequest) {
	for _, rm := range request.GetResourceMetrics() {
		if len(rm.ScopeMetrics) > 0 {
			s.metrics = append(s.metrics, rm.ScopeMetrics[0].Metrics...)
		}
	}
}

// GetMetrics returns the stored metrics.
func (s *MetricsStorage) GetMetrics() []*metricpb.Metric {
	// copy in order to not change.
	m := make([]*metricpb.Metric, 0, len(s.metrics))
	return append(m, s.metrics...)
}
