/*
 * Copyright (c) The go-kit Authors
 */

package grpcclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"

	"github.com/chief-of-state/cos-go-binding/grpcinterceptors"
	"github.com/chief-of-state/cos-go-binding/logging"
	"github.com/chief-of-state/cos-go-binding/requestid"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

// ConnectionBuilder is a builder to create GRPC connection to the GRPC Server
type ConnectionBuilder interface {
	WithOptions(opts ...grpc.DialOption)
	WithInsecure()
	WithUnaryInterceptors(interceptors []grpc.UnaryClientInterceptor)
	WithStreamInterceptors(interceptors []grpc.StreamClientInterceptor)
	WithKeepAliveParams(params keepalive.ClientParameters)
	GetConn(ctx context.Context, addr string) (*grpc.ClientConn, error)
	GetTLSConn(ctx context.Context, addr string) (*grpc.ClientConn, error)
}

// Builder is grpc client builder
type Builder struct {
	options              []grpc.DialOption
	transportCredentials credentials.TransportCredentials
}

// NewBuilder creates an instance of Builder
func NewBuilder() *Builder {
	return &Builder{}
}

// WithOptions set dial options
func (b *Builder) WithOptions(opts ...grpc.DialOption) *Builder {
	b.options = append(b.options, opts...)
	return b
}

// WithInsecure set the connection as insecure
func (b *Builder) WithInsecure() *Builder {
	b.options = append(b.options, grpc.WithInsecure()) // nolint
	return b
}

// WithBlock the dialing blocks until the  underlying connection is up.
// Without this, Dial returns immediately and connecting the server happens in background.
func (b *Builder) WithBlock() *Builder {
	b.options = append(b.options, grpc.WithBlock())
	return b
}

// WithKeepAliveParams set the keep alive params
// ClientParameters is used to set keepalive parameters on the client-side.
// These configure how the client will actively probe to notice when a
// connection is broken and send pings so intermediaries will be aware of the
// liveness of the connection. Make sure these parameters are set in
// coordination with the keepalive policy on the server, as incompatible
// settings can result in closing of connection.
func (b *Builder) WithKeepAliveParams(params keepalive.ClientParameters) *Builder {
	keepAlive := grpc.WithKeepaliveParams(params)
	b.options = append(b.options, keepAlive)
	return b
}

// WithUnaryInterceptors set a list of interceptors to the Grpc client for unary connection
// By default, gRPC doesn't allow one to have more than one interceptor either on the client nor on the server side.
// By using `grpc_middleware` we are able to provides convenient method to add a list of interceptors
func (b *Builder) WithUnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor) *Builder {
	b.options = append(b.options, grpc.WithUnaryInterceptor(grpcMiddleware.ChainUnaryClient(interceptors...)))
	return b
}

// WithStreamInterceptors set a list of interceptors to the Grpc client for stream connection
// By default, gRPC doesn't allow one to have more than one interceptor either on the client nor on the server side.
// By using `grpc_middleware` we are able to provides convenient method to add a list of interceptors
func (b *Builder) WithStreamInterceptors(interceptors ...grpc.StreamClientInterceptor) *Builder {
	b.options = append(b.options, grpc.WithStreamInterceptor(grpcMiddleware.ChainStreamClient(interceptors...)))
	return b
}

// WithClientTransportCredentials builds transport credentials for a gRPC client using the given properties.
func (b *Builder) WithClientTransportCredentials(insecureSkipVerify bool, certPool *x509.CertPool) *Builder {
	var tlsConf tls.Config

	if insecureSkipVerify {
		tlsConf.InsecureSkipVerify = true
		b.transportCredentials = credentials.NewTLS(&tlsConf)
		return b
	}

	tlsConf.RootCAs = certPool
	b.transportCredentials = credentials.NewTLS(&tlsConf)
	return b
}

// WithDefaultUnaryInterceptors sets the default unary interceptors for the grpc server
func (b *Builder) WithDefaultUnaryInterceptors() *Builder {
	return b.WithUnaryInterceptors(
		requestid.NewUnaryClientInterceptor(),
		grpcinterceptors.NewTracingClientUnaryInterceptor(),
		grpcinterceptors.NewClientMetricUnaryInterceptor(),
	)
}

// WithDefaultStreamInterceptors sets the default stream interceptors for the grpc server
func (b *Builder) WithDefaultStreamInterceptors() *Builder {
	return b.WithStreamInterceptors(
		requestid.NewStreamClientInterceptor(),
		grpcinterceptors.NewTracingClientStreamInterceptor(),
		grpcinterceptors.NewClientMetricStreamInterceptor(),
	)
}

// GetConn returns the client connection to the server
func (b *Builder) GetConn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	if addr == "" {
		return nil, fmt.Errorf("target connection parameter missing. address = %s", addr)
	}
	logging.Debugf("Target to connect = %s", addr)
	cc, err := grpc.DialContext(ctx, addr, b.options...)

	if err != nil {
		return nil, fmt.Errorf("unable to connect to client. address = %s. error = %+v", addr, err)
	}
	return cc, nil
}

// GetTLSConn returns client connection to the server
func (b *Builder) GetTLSConn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	b.options = append(b.options, grpc.WithTransportCredentials(b.transportCredentials))
	cc, err := grpc.DialContext(
		ctx,
		addr,
		b.options...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get tls conn. Unable to connect to client. address = %s: %w", addr, err)
	}
	return cc, nil
}

// GetClientConn return a grpc client connection
func GetClientConn(ctx context.Context, addr string) *grpc.ClientConn {
	// create the client builder
	clientBuilder := NewBuilder().
		WithDefaultUnaryInterceptors().
		WithDefaultStreamInterceptors().
		WithInsecure().
		WithKeepAliveParams(keepalive.ClientParameters{
			Time:                1200 * time.Second,
			PermitWithoutStream: true,
		})
	// get the gRPC client connection
	conn, err := clientBuilder.GetConn(ctx, addr)
	// handle the connection error
	if err != nil {
		logging.Fatalf("failed to create grpc service client, %s", err.Error())
	}
	// return the client connection created
	return conn
}
