/*
 * Copyright (c) The go-kit Authors
 */

package grpcinterceptors

import (
	"github.com/chief-of-state/cos-go-binding/logging"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewRecoveryUnaryInterceptor recovers from an unexpected panic
// Recovery handlers should typically be last in the chain so that other middleware
// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
func NewRecoveryUnaryInterceptor() grpc.UnaryServerInterceptor {
	// Define customfunc to handle panic
	customFunc := func(p interface{}) (err error) {
		logging.Error(p)
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpcRecovery.Option{
		grpcRecovery.WithRecoveryHandler(customFunc),
	}

	return grpcRecovery.UnaryServerInterceptor(opts...)
}

// NewRecoveryStreamInterceptor recovers from an unexpected panic
// Recovery handlers should typically be last in the chain so that other middleware
// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
func NewRecoveryStreamInterceptor() grpc.StreamServerInterceptor {
	// Define customfunc to handle panic
	customFunc := func(p interface{}) (err error) {
		logging.Error(p)
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpcRecovery.Option{
		grpcRecovery.WithRecoveryHandler(customFunc),
	}

	return grpcRecovery.StreamServerInterceptor(opts...)
}
