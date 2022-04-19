/*
 * Copyright (c) The go-kit Authors
 */

package grpcclient

import (
	"context"
	"testing"

	helloworldv1 "github.com/chief-of-state/cos-go-binding/gen/helloworld/v1"
	"github.com/chief-of-state/cos-go-binding/grpctesting"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type ClientTestSuite struct {
	suite.Suite

	server     grpctesting.InProcessServer
	clientConn *grpc.ClientConn
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

// SetupTest will run before each test in the suite.
func (s *ClientTestSuite) SetupTest() {
	builder := grpctesting.NewInProcessServerBuilder()
	s.server = builder.Build()
	s.server.RegisterService(func(server *grpc.Server) {
		helloworldv1.RegisterGreeterServer(server, &grpctesting.MockedService{})
	})
	err := s.server.Start()
	s.Assert().NoError(err)
}

// TearDownTest will run after each test in the suite
func (s *ClientTestSuite) TearDownTest() {
	s.server.Cleanup()
	err := s.clientConn.Close()
	s.Assert().NoError(err)
}

func (s *ClientTestSuite) TestSayHello() {
	s.Run("with context", func() {
		ctx := context.Background()
		var err error
		clientBuilder := NewBuilder().
			WithInsecure().
			WithDefaultStreamInterceptors().
			WithDefaultUnaryInterceptors().
			WithBlock().
			WithOptions(grpc.WithContextDialer(grpctesting.GetBufDialer(s.server.GetListener())))

		s.clientConn, err = clientBuilder.GetConn(ctx, "localhost:50051")

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
	})
}
