/*
 * Copyright (c) The go-kit Authors
 */

package grpctesting

import (
	"context"
	"testing"

	helloworldv1 "github.com/chief-of-state/cos-go-binding/gen/helloworld/v1"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type InProcessTestSuite struct {
	suite.Suite

	server     InProcessServer
	clientConn *grpc.ClientConn
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestInProcessTestSuite(t *testing.T) {
	suite.Run(t, new(InProcessTestSuite))
}

func (s *InProcessTestSuite) SetupTest() {
	builder := InProcessServerBuilder{}
	s.server = builder.Build()
	s.server.RegisterService(func(server *grpc.Server) {
		helloworldv1.RegisterGreeterServer(server, &MockedService{})
	})
	err := s.server.Start()
	s.Assert().NoError(err)
}

func (s *InProcessTestSuite) TearDownSuite() {
	s.server.Cleanup()
	err := s.clientConn.Close()
	s.Assert().NoError(err)
}

// TestSayHello will test the HelloWorld service using an in-memory data transfer
// instead of the normal networking counterpart
func (s *InProcessTestSuite) TestSayHello() {
	ctx := context.Background()
	var err error
	s.clientConn, err = GetInProcessClientConn(ctx, s.server.GetListener(), []grpc.DialOption{})
	if err != nil {
		s.T().Fatalf("Failed to dial bufnet: %v", err)
	}

	client := helloworldv1.NewGreeterClient(s.clientConn)
	request := &helloworldv1.HelloRequest{Name: "test"}
	resp, err := client.SayHello(ctx, request)
	if err != nil {
		s.T().Fatalf("SayHello failed: %v", err)
	}
	s.Assert().Equal(resp.Message, "This is a mocked service test")
}
