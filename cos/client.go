package cos

import (
	"context"

	cospb "github.com/chief-of-state/cos-go-sdk/cospb/chief_of_state/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ClientIFace is the interface for a cos client
type ClientIFace[T proto.Message] interface {
	// ProcessCommand sends a command to COS and returns the resulting state as T and metadata.
	// The returned error can either be a gRPC error or casting error
	// gRPC status codes: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	ProcessCommand(ctx context.Context, entityID string, command proto.Message) (T, *cospb.MetaData, error)
	// GetState retrieves the current state as T of an entity and its metadata
	// The returned error can either be a gRPC error or casting error
	// gRPC status codes: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	GetState(ctx context.Context, entityID string) (T, *cospb.MetaData, error)
}

// CosClient implements the Client interface
type CosClient[T proto.Message] struct {
	Remote cospb.ChiefOfStateServiceClient
}

// NewClient creates a new instance of Client
func NewClient[T proto.Message](conn *grpc.ClientConn) CosClient[T] {
	return CosClient[T]{
		Remote: cospb.NewChiefOfStateServiceClient(conn),
	}
}

// ProcessCommand sends a command to COS and returns the resulting state as T and metadata.
// The returned error can either be a gRPC error or casting error
// gRPC status codes: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
func (c CosClient[T]) ProcessCommand(ctx context.Context, entityID string, command proto.Message) (T, *cospb.MetaData, error) {
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
	response, err := c.Remote.ProcessCommand(ctx, request)
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

// GetState retrieves the current state as T of an entity and its metadata
// The returned error can either be a gRPC error or casting error
// gRPC status codes: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
func (c CosClient[T]) GetState(ctx context.Context, entityID string) (T, *cospb.MetaData, error) {
	var defaultT T
	// call CoS
	response, err := c.Remote.GetState(ctx, &cospb.GetStateRequest{EntityId: entityID})
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

// unpackState takes an any to unpack into T
// The returned error can be a casting error
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
