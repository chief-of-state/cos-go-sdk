/*
 * Copyright (c) The go-kit Authors
 */

package grpctesting

import (
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
	v1 "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	metricpb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type MockCollectorTestSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMockCollectorTestSuite(t *testing.T) {
	suite.Run(t, new(MockCollectorTestSuite))
}

func (s *MockCollectorTestSuite) TestMakeMockCollector() {
	mockCollector := MakeMockCollector(&MockConfig{Endpoint: "localhost:0"})
	s.Assert().NotNil(mockCollector)
}

func (s *MockCollectorTestSuite) TestGetMetrics() {
	mockCollector := MakeMockCollector(&MockConfig{Endpoint: "localhost:0"})
	s.Assert().NotNil(mockCollector)
	metrics := mockCollector.GetMetrics()
	s.Assert().NotNil(metrics)
	s.Assert().Empty(metrics)
}

func (s *MockCollectorTestSuite) TestGetEndPoint() {
	mockCollector := MakeMockCollector(&MockConfig{Endpoint: "localhost:4774"})
	s.Assert().NotNil(mockCollector)
	endpoint := mockCollector.GetEndPoint()
	s.Assert().Empty(endpoint)
}

func (s *MockCollectorTestSuite) TestRunMockCollector() {
	mockCollector, err := RunMockCollector()
	s.Assert().NoError(err)
	s.Assert().NotNil(mockCollector)
	endpoint := mockCollector.GetEndPoint()
	s.Assert().NotEmpty(endpoint)
	s.Assert().Contains(endpoint, "127.0.0.1")
	err = mockCollector.Stop()
	s.Assert().NoError(err)
}

func (s *MockCollectorTestSuite) TestRunMockCollectorAtEndpoint() {
	mockCollector, err := RunMockCollectorAtEndpoint("127.0.0.1:4447")
	s.Assert().NoError(err)
	s.Assert().NotNil(mockCollector)
	endpoint := mockCollector.GetEndPoint()
	s.Assert().NotEmpty(endpoint)
	s.Assert().Equal("127.0.0.1:4447", endpoint)
	err = mockCollector.Stop()
	s.Assert().NoError(err)
}

func (s *MockCollectorTestSuite) TestRunMockCollectorWithConfig() {
	s.Run("valid endpoint", func() {
		mockCollector, err := RunMockCollectorWithConfig(&MockConfig{
			Endpoint: "127.0.0.1:4447",
		})
		s.Assert().NoError(err)
		s.Assert().NotNil(mockCollector)
		endpoint := mockCollector.GetEndPoint()
		s.Assert().NotEmpty(endpoint)
		s.Assert().Equal("127.0.0.1:4447", endpoint)
		err = mockCollector.Stop()
		s.Assert().NoError(err)
	})

	s.Run("invalid endpoint", func() {
		mockCollector, err := RunMockCollectorWithConfig(&MockConfig{
			Endpoint: "some-point",
		})
		s.Assert().Error(err)
		s.Assert().Nil(mockCollector)
	})
}

func (s *MockCollectorTestSuite) TestAddMetrics() {
	s.Run("when there are some metrics", func() {
		metricStorage := NewMetricsStorage()
		metricStorage.AddMetrics(&v1.ExportMetricsServiceRequest{
			ResourceMetrics: []*metricpb.ResourceMetrics{
				{
					ScopeMetrics: []*metricpb.ScopeMetrics{
						{
							Metrics: []*metricpb.Metric{
								{
									Name:        "metric-1",
									Description: "metric-1",
									Unit:        "unit-1",
									Data:        nil,
								},
							},
						},
					},
				},
			},
		})

		s.Assert().NotEmpty(metricStorage.metrics)
		s.Assert().Equal(1, len(metricStorage.metrics))
	})

	s.Run("when there are no metrics", func() {
		metricStorage := NewMetricsStorage()
		metricStorage.AddMetrics(&v1.ExportMetricsServiceRequest{
			ResourceMetrics: []*metricpb.ResourceMetrics{},
		})

		s.Assert().Empty(metricStorage.metrics)
	})
}

func (s *MockCollectorTestSuite) TestStorageGetMetrics() {
	s.Run("when there no metrics", func() {
		ms := NewMetricsStorage()
		metrics := ms.GetMetrics()
		s.Assert().Empty(metrics)
	})

	s.Run("when there some metrics", func() {
		ms := NewMetricsStorage()
		ms.metrics = []*metricpb.Metric{
			{
				Name:        "metric-1",
				Description: "metric-1",
				Unit:        "unit-1",
				Data:        nil,
			},
			{
				Name:        "metric-2",
				Description: "metric-2",
				Unit:        "unit-2",
				Data:        nil,
			},
		}

		metrics := ms.GetMetrics()
		s.Assert().NotEmpty(metrics)
		s.Assert().Equal(2, len(metrics))
	})
}

func (s *MockCollectorTestSuite) TestListener() {
	ln, err := net.Listen("tcp", "localhost:50051")
	s.Assert().NoError(err)

	listnr := newListener(ln)
	s.Assert().NotNil(listnr)

	addr := listnr.Addr()
	s.Assert().NotNil(addr)

	err = listnr.Close()
	s.Assert().NoError(err)
}
