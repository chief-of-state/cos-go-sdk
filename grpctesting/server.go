/*
 * Copyright (c) The go-kit Authors
 */

// Package grpctesting In Processing server uses memory to transfer data between the server and the client
// This is ideal for testing purpose as it not require networking to run integration tests
package grpctesting

import (
	"crypto/tls"

	"github.com/chief-of-state/cos-go-binding/logging"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/test/bufconn"
)

// InProcessServer server interface
type InProcessServer interface {
	Start() error
	RegisterService(reg func(*grpc.Server))
	Cleanup()
	GetListener() *bufconn.Listener
}

// InProcessServerBuilder in-processing grpc server builder
type InProcessServerBuilder struct {
	options []grpc.ServerOption
}

// NewInProcessServerBuilder creates an instance of InProcessServerBuilder
func NewInProcessServerBuilder() *InProcessServerBuilder {
	return new(InProcessServerBuilder)
}

// WithOption configures how we set up the connection.
func (sb *InProcessServerBuilder) WithOption(o grpc.ServerOption) *InProcessServerBuilder {
	sb.options = append(sb.options, o)
	return sb
}

// WithStreamInterceptors set a list of interceptors to the Grpc server for stream connection
// By default, gRPC doesn't allow one to have more than one interceptor either on the client nor on the server side.
// By using `grpcMiddleware` we are able to provides convenient method to add a list of interceptors
func (sb *InProcessServerBuilder) WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) *InProcessServerBuilder {
	chain := grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(interceptors...))
	sb.WithOption(chain)
	return sb
}

// WithUnaryInterceptors set a list of interceptors to the Grpc server for unary connection
// By default, gRPC doesn't allow one to have more than one interceptor either on the client nor on the server side.
// By using `grpcMiddleware` we are able to provides convenient method to add a list of interceptors
func (sb *InProcessServerBuilder) WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) *InProcessServerBuilder {
	chain := grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(interceptors...))
	sb.WithOption(chain)
	return sb
}

// WithTLSCert sets credentials for server connections
func (sb *InProcessServerBuilder) WithTLSCert(cert *tls.Certificate) *InProcessServerBuilder {
	sb.WithOption(grpc.Creds(credentials.NewServerTLSFromCert(cert)))
	return sb
}

//Build is responsible for building a Fiji GRPC server
func (sb *InProcessServerBuilder) Build() InProcessServer {
	server, listener := GetInProcessServer(sb.options)
	return &grpcServer{server, listener}
}

type grpcServer struct {
	server   *grpc.Server
	listener *bufconn.Listener
}

// GetListener register the services to the server
func (s *grpcServer) GetListener() *bufconn.Listener {
	return s.listener
}

// RegisterService register the services to the server
func (s *grpcServer) RegisterService(reg func(*grpc.Server)) {
	reg(s.server)
}

// Start the GRPC server
func (s *grpcServer) Start() error {
	go s.serv()
	logging.Printf("In processing server started")
	return nil
}

// Cleanup stops the server and close the tcp listener
func (s *grpcServer) Cleanup() {
	s.server.Stop()
	_ = s.listener.Close()
	logging.Print("Server stopped")
}

func (s *grpcServer) serv() {
	if err := s.server.Serve(s.listener); err != nil {
		logging.Fatalf("failed to serve: %+v", err)
	}
}
