package cos

import (
	"context"
	"net"
	"testing"

	mocks "github.com/chief-of-state/cos-go-sdk/cosmocks"
	cospb "github.com/chief-of-state/cos-go-sdk/cospb/chief_of_state/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type clientSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestClient(t *testing.T) {
	suite.Run(t, new(clientSuite))
}

func (s *clientSuite) TestProcessCommand() {

	s.Run("with nil command", func() {
		// create the remote client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		state, meta, err := mockCos.ProcessCommand(context.TODO(), uuid.NewString(), nil)
		expectedError := status.Error(codes.Internal, "command is missing")
		s.Assert().Nil(state)
		s.Assert().Nil(meta)
		s.Assert().EqualError(err, expectedError.Error())
	})
	s.Run("with happy path", func() {
		ctx := context.TODO()
		// create the various ID
		now := timestamppb.Now()
		entityID := "foo"
		// create the current state
		currentState := &wrapperspb.StringValue{}
		anypbState, err := anypb.New(currentState)
		s.Assert().NoError(err)
		cosMeta := &cospb.MetaData{
			EntityId:       entityID,
			RevisionNumber: 1,
			RevisionDate:   now,
		}
		// create the process command response
		cosResp := &cospb.ProcessCommandResponse{State: anypbState, Meta: cosMeta}
		// create the remote client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockRemoteClient.On("ProcessCommand", ctx, mock.Anything).Return(cosResp, nil)
		// create the CoS client
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		cmd := &wrapperspb.StringValue{}
		state, meta, err := mockCos.ProcessCommand(ctx, entityID, cmd)
		s.Assert().NoError(err)
		s.Assert().NotNil(meta)
		s.Assert().NotNil(state)
		s.Assert().True(proto.Equal(currentState, state))
		s.Assert().True(proto.Equal(cosMeta, meta))
		mockRemoteClient.AssertExpectations(s.T())
	})
	s.Run("with remote client failure", func() {
		ctx := context.TODO()
		id := uuid.NewString()

		// create the remote client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockRemoteClient.On("ProcessCommand", ctx, mock.Anything).Return(nil, status.Error(codes.Internal, ""))
		// create the CoS client
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		cmd := &wrapperspb.StringValue{}
		state, meta, err := mockCos.ProcessCommand(ctx, id, cmd)
		s.Assert().Error(err)
		s.Assert().Nil(meta)
		s.Assert().Nil(state)
		mockRemoteClient.AssertExpectations(s.T())
	})
	s.Run("with invalid state returned", func() {
		ctx := context.TODO()
		// create the various ID
		communityID := uuid.NewString()
		now := timestamppb.Now()
		anypbState, err := anypb.New(wrapperspb.Int32(100))
		s.Assert().NoError(err)
		cosMeta := &cospb.MetaData{
			EntityId:       communityID,
			RevisionNumber: 2,
			RevisionDate:   now,
		}
		// create the process command response
		cosResp := &cospb.ProcessCommandResponse{State: anypbState, Meta: cosMeta}
		// create the remote client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockRemoteClient.On("ProcessCommand", ctx, mock.Anything).Return(cosResp, nil)
		// create the CoS client
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		cmd := &wrapperspb.StringValue{}
		state, meta, err := mockCos.ProcessCommand(ctx, communityID, cmd)
		s.Assert().Error(err)
		s.Assert().Nil(meta)
		s.Assert().Nil(state)
		mockRemoteClient.AssertExpectations(s.T())
	})
}
func (s *clientSuite) TestGetState() {
	s.Run("with happy path", func() {
		ctx := context.TODO()
		id := uuid.NewString()
		now := timestamppb.Now()
		// create the current state
		currentState := &wrapperspb.StringValue{}
		anypbState, err := anypb.New(currentState)
		s.Assert().NoError(err)
		s.Assert().NotNil(anypbState)
		cosMeta := &cospb.MetaData{
			EntityId:       id,
			RevisionNumber: 2,
			RevisionDate:   now,
		}
		// create the process command response
		cosResp := &cospb.GetStateResponse{State: anypbState, Meta: cosMeta}
		// create the client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		mockRemoteClient.On("GetState", ctx, mock.Anything).Return(cosResp, nil)
		state, meta, err := mockCos.GetState(ctx, id)
		s.Assert().NoError(err)
		s.Assert().NotNil(meta)
		s.Assert().NotNil(state)
		s.Assert().True(proto.Equal(currentState, state))
		s.Assert().True(proto.Equal(cosMeta, meta))
		mockRemoteClient.AssertExpectations(s.T())
	})
	s.Run("with CoS failure", func() {
		ctx := context.TODO()
		id := uuid.NewString()
		// create the current state
		currentState := &wrapperspb.StringValue{}
		anypbState, err := anypb.New(currentState)
		s.Assert().NoError(err)
		s.Assert().NotNil(anypbState)
		// create the remote client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockRemoteClient.On("GetState", ctx, mock.Anything).Return(nil, status.Error(codes.Unavailable, ""))
		// create the CoS client
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		state, meta, err := mockCos.GetState(ctx, id)
		s.Assert().Error(err)
		s.Assert().Nil(meta)
		s.Assert().Nil(state)
		mockRemoteClient.AssertExpectations(s.T())
	})
	s.Run("with invalid state", func() {
		ctx := context.TODO()
		id := uuid.NewString()
		now := timestamppb.Now()
		// create the current state
		anypbState, err := anypb.New(wrapperspb.Int32(100))
		s.Assert().NoError(err)
		s.Assert().NotNil(anypbState)
		s.Assert().NoError(err)
		s.Assert().NotNil(anypbState)
		cosMeta := &cospb.MetaData{
			EntityId:       id,
			RevisionNumber: 2,
			RevisionDate:   now,
		}
		// create the process command response
		cosResp := &cospb.GetStateResponse{State: anypbState, Meta: cosMeta}
		// create the remote client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockRemoteClient.On("GetState", ctx, mock.Anything).Return(cosResp, nil)
		// create the CoS client
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		state, meta, err := mockCos.GetState(ctx, id)
		s.Assert().Error(err)
		s.Assert().Nil(meta)
		s.Assert().Nil(state)
		mockRemoteClient.AssertExpectations(s.T())
	})
	s.Run("with not found", func() {
		ctx := context.TODO()
		id := uuid.NewString()
		// create the remote client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockRemoteClient.On("GetState", ctx, mock.Anything).Return(nil, status.Error(codes.NotFound, "state not found"))
		// create the CoS client
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		state, meta, err := mockCos.GetState(ctx, id)
		s.Assert().NoError(err)
		s.Assert().Nil(meta)
		s.Assert().Nil(state)
		mockRemoteClient.AssertExpectations(s.T())
	})
	s.Run("with nil response", func() {
		ctx := context.TODO()
		id := uuid.NewString()
		// create the remote client
		mockRemoteClient := &mocks.ChiefOfStateServiceClient{}
		mockRemoteClient.On("GetState", ctx, mock.Anything).Return(nil, nil)
		// create the CoS client
		mockCos := CosClient[*wrapperspb.StringValue]{Remote: mockRemoteClient}
		state, meta, err := mockCos.GetState(ctx, id)
		s.Assert().NoError(err)
		s.Assert().Nil(meta)
		s.Assert().Nil(state)
		mockRemoteClient.AssertExpectations(s.T())
	})
}
func (s *clientSuite) TestNewClient() {
	s.Run("happy path", func() {
		// create a context
		ctx := context.TODO()
		listen := bufconn.Listen(1024 * 1024)
		target := "bufnet"
		grpcClient, err := grpc.DialContext(
			ctx,
			target,
			grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return listen.Dial() }),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		s.Assert().NoError(err)
		// this will work because grpc connection won't wait for connections to be
		// established, and connecting happens in the background
		cosClient := NewClient[*wrapperspb.StringValue](grpcClient)
		s.Assert().NotNil(cosClient)
	})
}
func (s *clientSuite) TestUnpackState() {
	s.Run("happy path with Poll State", func() {
		// create a new state
		state := &wrapperspb.StringValue{}
		// pack that state into any
		any, err := anypb.New(state)
		s.Assert().NoError(err)
		s.Assert().NotNil(any)

		unpacked, err := unpackState[*wrapperspb.StringValue](any)
		s.Assert().NoError(err)
		s.Assert().True(proto.Equal(state, unpacked))
	})
	s.Run("happy path with Vote State", func() {
		// create a new state
		state := &wrapperspb.StringValue{}
		// pack that state into any
		any, err := anypb.New(state)
		s.Assert().NoError(err)
		s.Assert().NotNil(any)

		unpacked, err := unpackState[*wrapperspb.StringValue](any)
		s.Assert().NoError(err)
		s.Assert().True(proto.Equal(state, unpacked))
	})
	s.Run("with empty proto", func() {
		// create an empty proto message
		empty := new(emptypb.Empty)
		// pack into any
		any, err := anypb.New(empty)
		s.Assert().NoError(err)
		s.Assert().NotNil(any)
		unpacked, err := unpackState[*wrapperspb.StringValue](any)
		s.Assert().NoError(err)
		s.Assert().Nil(unpacked)
	})
	s.Run("with invalid state", func() {
		// create a wrong state
		any, err := anypb.New(wrapperspb.Int32(100))
		s.Assert().NoError(err)
		s.Assert().NotNil(any)
		unpacked, err := unpackState[*wrapperspb.StringValue](any)
		s.Assert().Error(err)
		s.Assert().Nil(unpacked)
	})
	s.Run("with invalid any state", func() {
		// create a wrong state
		any := &anypb.Any{
			TypeUrl: "",
			Value:   nil,
		}
		unpacked, err := unpackState[*wrapperspb.StringValue](any)
		s.Assert().Error(err)
		s.Assert().Nil(unpacked)
	})
}
