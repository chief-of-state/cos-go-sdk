/*
 * Copyright (c) The go-kit Authors
 */

package grpctesting

import (
	"context"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// InProcessClientBuilder in-processing grpc client builder
type InProcessClientBuilder struct {
	Server  InProcessServer
	options []grpc.DialOption
	ctx     context.Context
}

// NewInProcessClientBuilder creates a new instance of InProcessClientBuilder
func NewInProcessClientBuilder(server InProcessServer) *InProcessClientBuilder {
	return &InProcessClientBuilder{
		Server: server,
	}
}

// WithContext set the context to be used in the dial
func (b *InProcessClientBuilder) WithContext(ctx context.Context) *InProcessClientBuilder {
	b.ctx = ctx
	return b
}

// WithOptions set dial options
func (b *InProcessClientBuilder) WithOptions(opts ...grpc.DialOption) *InProcessClientBuilder {
	b.options = append(b.options, opts...)
	return b
}

// WithInsecure set the connection as insecure
func (b *InProcessClientBuilder) WithInsecure() *InProcessClientBuilder {
	b.options = append(b.options, grpc.WithInsecure()) // nolint
	return b
}

// WithUnaryInterceptors set a list of interceptors to the Grpc client for unary connection
// By default, gRPC doesn't allow one to have more than one interceptor either on the client nor on the Server side.
// By using `grpcMiddleware` we are able to provides convenient method to add a list of interceptors
func (b *InProcessClientBuilder) WithUnaryInterceptors(interceptors []grpc.UnaryClientInterceptor) *InProcessClientBuilder {
	b.options = append(b.options, grpc.WithUnaryInterceptor(grpcMiddleware.ChainUnaryClient(interceptors...)))
	return b
}

// WithStreamInterceptors set a list of interceptors to the Grpc client for stream connection
// By default, gRPC doesn't allow one to have more than one interceptor either on the client nor on the Server side.
// By using `grpcMiddleware` we are able to provides convenient method to add a list of interceptors
func (b *InProcessClientBuilder) WithStreamInterceptors(interceptors []grpc.StreamClientInterceptor) *InProcessClientBuilder {
	b.options = append(b.options, grpc.WithStreamInterceptor(grpcMiddleware.ChainStreamClient(interceptors...)))
	return b
}

// GetConn returns the client connection to the Server
func (b *InProcessClientBuilder) GetConn(addr string, port string) (*grpc.ClientConn, error) {
	ctx := b.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	return GetInProcessClientConn(ctx, b.Server.GetListener(), b.options)
}
