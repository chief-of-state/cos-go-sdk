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

type ClientTestSuite struct {
	suite.Suite

	server     InProcessServer
	clientConn *grpc.ClientConn
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

// SetupTest will run before each test in the suite.
func (s *ClientTestSuite) SetupTest() {
	builder := InProcessServerBuilder{}
	s.server = builder.Build()
	s.server.RegisterService(func(server *grpc.Server) {
		helloworldv1.RegisterGreeterServer(server, &MockedService{})
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
		clientBuilder := NewInProcessClientBuilder(s.server).
			WithInsecure().
			WithContext(ctx).
			WithOptions(grpc.WithContextDialer(GetBufDialer(s.server.GetListener())))

		s.clientConn, err = clientBuilder.GetConn("localhost", "50051")
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

	s.Run("without context", func() {
		ctx := context.Background()
		var err error
		clientBuilder := NewInProcessClientBuilder(s.server).
			WithInsecure().
			WithOptions(grpc.WithContextDialer(GetBufDialer(s.server.GetListener())))

		s.clientConn, err = clientBuilder.GetConn("localhost", "50051")
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
