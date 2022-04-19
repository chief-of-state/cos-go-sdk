package cos

import (
	"context"
	"fmt"

	cospb "github.com/chief-of-state/cos-go-binding/gen/chief_of_state/v1"
	"github.com/chief-of-state/cos-go-binding/grpcclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Client is used by both the service and the consumer implementation.
type Client[T proto.Message] interface {
	ProcessCommand(ctx context.Context, entityID string, command proto.Message) (T, *cospb.MetaData, error)
	GetState(ctx context.Context, entityID string) (T, *cospb.MetaData, error)
}

// cosClient implements the Client interface
type cosClient[T proto.Message] struct {
	remote cospb.ChiefOfStateServiceClient
}

// NewClient creates a new instance of Client
func NewClient[T proto.Message](ctx context.Context, cosHost string, cosPort int) (Client[T], error) {
	conn := grpcclient.GetClientConn(ctx, fmt.Sprintf("%v:%v", cosHost, cosPort))
	return cosClient[T]{
		remote: cospb.NewChiefOfStateServiceClient(conn),
	}, nil
}

// ProcessCommand sends a command to COS and returns the resulting state and metadata
func (c cosClient[T]) ProcessCommand(ctx context.Context, entityID string, command proto.Message) (T, *cospb.MetaData, error) {
	var defaultT T
	// require a command
	if command == nil {
		return defaultT, nil, status.Error(codes.Internal, "command is missing")
	}

	// pack command into Any
	cmdAny, _ := anypb.New(command)

	// construct COS request
	request := &cospb.ProcessCommandRequest{
		EntityId: entityID,
		Command:  cmdAny,
	}

	// call COS get response
	response, err := c.remote.ProcessCommand(ctx, request)
	if err != nil {
		return defaultT, nil, err
	}

	// unpack the resulting state
	resultingState, err := unpackState[T](response.GetState())
	if err != nil {
		return defaultT, nil, err
	}

	// return the company and the metadata
	return resultingState, response.GetMeta(), nil
}

// GetState retrieves the current  state of an entity and its metadata
func (c cosClient[T]) GetState(ctx context.Context, entityID string) (T, *cospb.MetaData, error) {
	var defaultT T
	// call CoS
	response, err := c.remote.GetState(ctx, &cospb.GetStateRequest{EntityId: entityID})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.NotFound {
				return defaultT, nil, nil
			}
		}

		return defaultT, nil, err
	}

	// handle nil response like a NOT_FOUND
	if response == nil {
		return defaultT, nil, nil
	}

	// unpack the resulting state
	resultingState, err := unpackState[T](response.GetState())
	if err != nil {
		return defaultT, nil, err
	}

	// return
	return resultingState, response.GetMeta(), nil
}

// unpackState takes an any to unpack into either pollsv1.Poll or pollsv1.Vote
func unpackState[T proto.Message](any *anypb.Any) (T, error) {
	var defaultT T

	msg, err := any.UnmarshalNew()
	if err != nil {
		return defaultT, err
	}

	switch v := msg.(type) {
	case T:
		return v, nil
	case *emptypb.Empty:
		return defaultT, nil
	default:
		return defaultT, status.Errorf(codes.Internal, "got %s", any.GetTypeUrl())
	}
}
