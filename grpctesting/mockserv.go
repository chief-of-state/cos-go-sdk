/*
 * Copyright (c) The go-kit Authors
 */

package grpctesting

import (
	"context"

	"google.golang.org/grpc"

	helloworldv1 "github.com/chief-of-state/cos-go-binding/gen/helloworld/v1"
)

// MockedService is only used in grpc unit tests
type MockedService struct {
}

// SayHello will handle the HelloRequest and return the appropriate response
func (s *MockedService) SayHello(ctx context.Context, in *helloworldv1.HelloRequest) (*helloworldv1.HelloReply, error) {
	return &helloworldv1.HelloReply{Message: "This is a mocked service " + in.Name}, nil
}

func (s *MockedService) RegisterService(server *grpc.Server) {
	helloworldv1.RegisterGreeterServer(server, s)
}
