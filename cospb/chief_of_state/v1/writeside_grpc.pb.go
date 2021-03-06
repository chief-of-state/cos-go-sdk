// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: chief_of_state/v1/writeside.proto

package chief_of_statev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WriteSideHandlerServiceClient is the client API for WriteSideHandlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriteSideHandlerServiceClient interface {
	// Processes every command sent by ChiefOfState and return either a response
	// containing an event to persist or a simple reply.
	HandleCommand(ctx context.Context, in *HandleCommandRequest, opts ...grpc.CallOption) (*HandleCommandResponse, error)
	// Processes every event sent by ChiefOfState by applying the event to the
	// current state to return a new state.
	HandleEvent(ctx context.Context, in *HandleEventRequest, opts ...grpc.CallOption) (*HandleEventResponse, error)
}

type writeSideHandlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWriteSideHandlerServiceClient(cc grpc.ClientConnInterface) WriteSideHandlerServiceClient {
	return &writeSideHandlerServiceClient{cc}
}

func (c *writeSideHandlerServiceClient) HandleCommand(ctx context.Context, in *HandleCommandRequest, opts ...grpc.CallOption) (*HandleCommandResponse, error) {
	out := new(HandleCommandResponse)
	err := c.cc.Invoke(ctx, "/chief_of_state.v1.WriteSideHandlerService/HandleCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writeSideHandlerServiceClient) HandleEvent(ctx context.Context, in *HandleEventRequest, opts ...grpc.CallOption) (*HandleEventResponse, error) {
	out := new(HandleEventResponse)
	err := c.cc.Invoke(ctx, "/chief_of_state.v1.WriteSideHandlerService/HandleEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteSideHandlerServiceServer is the server API for WriteSideHandlerService service.
// All implementations should embed UnimplementedWriteSideHandlerServiceServer
// for forward compatibility
type WriteSideHandlerServiceServer interface {
	// Processes every command sent by ChiefOfState and return either a response
	// containing an event to persist or a simple reply.
	HandleCommand(context.Context, *HandleCommandRequest) (*HandleCommandResponse, error)
	// Processes every event sent by ChiefOfState by applying the event to the
	// current state to return a new state.
	HandleEvent(context.Context, *HandleEventRequest) (*HandleEventResponse, error)
}

// UnimplementedWriteSideHandlerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWriteSideHandlerServiceServer struct {
}

func (UnimplementedWriteSideHandlerServiceServer) HandleCommand(context.Context, *HandleCommandRequest) (*HandleCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCommand not implemented")
}
func (UnimplementedWriteSideHandlerServiceServer) HandleEvent(context.Context, *HandleEventRequest) (*HandleEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleEvent not implemented")
}

// UnsafeWriteSideHandlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriteSideHandlerServiceServer will
// result in compilation errors.
type UnsafeWriteSideHandlerServiceServer interface {
	mustEmbedUnimplementedWriteSideHandlerServiceServer()
}

func RegisterWriteSideHandlerServiceServer(s grpc.ServiceRegistrar, srv WriteSideHandlerServiceServer) {
	s.RegisterService(&WriteSideHandlerService_ServiceDesc, srv)
}

func _WriteSideHandlerService_HandleCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteSideHandlerServiceServer).HandleCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chief_of_state.v1.WriteSideHandlerService/HandleCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteSideHandlerServiceServer).HandleCommand(ctx, req.(*HandleCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriteSideHandlerService_HandleEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteSideHandlerServiceServer).HandleEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chief_of_state.v1.WriteSideHandlerService/HandleEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteSideHandlerServiceServer).HandleEvent(ctx, req.(*HandleEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WriteSideHandlerService_ServiceDesc is the grpc.ServiceDesc for WriteSideHandlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriteSideHandlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chief_of_state.v1.WriteSideHandlerService",
	HandlerType: (*WriteSideHandlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleCommand",
			Handler:    _WriteSideHandlerService_HandleCommand_Handler,
		},
		{
			MethodName: "HandleEvent",
			Handler:    _WriteSideHandlerService_HandleEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chief_of_state/v1/writeside.proto",
}
