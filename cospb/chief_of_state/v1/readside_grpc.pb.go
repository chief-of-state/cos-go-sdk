// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: chief_of_state/v1/readside.proto

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

// ReadSideHandlerServiceClient is the client API for ReadSideHandlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReadSideHandlerServiceClient interface {
	// Helps build a read model from persisted events and snpahots
	HandleReadSide(ctx context.Context, in *HandleReadSideRequest, opts ...grpc.CallOption) (*HandleReadSideResponse, error)
}

type readSideHandlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReadSideHandlerServiceClient(cc grpc.ClientConnInterface) ReadSideHandlerServiceClient {
	return &readSideHandlerServiceClient{cc}
}

func (c *readSideHandlerServiceClient) HandleReadSide(ctx context.Context, in *HandleReadSideRequest, opts ...grpc.CallOption) (*HandleReadSideResponse, error) {
	out := new(HandleReadSideResponse)
	err := c.cc.Invoke(ctx, "/chief_of_state.v1.ReadSideHandlerService/HandleReadSide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadSideHandlerServiceServer is the server API for ReadSideHandlerService service.
// All implementations should embed UnimplementedReadSideHandlerServiceServer
// for forward compatibility
type ReadSideHandlerServiceServer interface {
	// Helps build a read model from persisted events and snpahots
	HandleReadSide(context.Context, *HandleReadSideRequest) (*HandleReadSideResponse, error)
}

// UnimplementedReadSideHandlerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedReadSideHandlerServiceServer struct {
}

func (UnimplementedReadSideHandlerServiceServer) HandleReadSide(context.Context, *HandleReadSideRequest) (*HandleReadSideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleReadSide not implemented")
}

// UnsafeReadSideHandlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReadSideHandlerServiceServer will
// result in compilation errors.
type UnsafeReadSideHandlerServiceServer interface {
	mustEmbedUnimplementedReadSideHandlerServiceServer()
}

func RegisterReadSideHandlerServiceServer(s grpc.ServiceRegistrar, srv ReadSideHandlerServiceServer) {
	s.RegisterService(&ReadSideHandlerService_ServiceDesc, srv)
}

func _ReadSideHandlerService_HandleReadSide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleReadSideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadSideHandlerServiceServer).HandleReadSide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chief_of_state.v1.ReadSideHandlerService/HandleReadSide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadSideHandlerServiceServer).HandleReadSide(ctx, req.(*HandleReadSideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReadSideHandlerService_ServiceDesc is the grpc.ServiceDesc for ReadSideHandlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReadSideHandlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chief_of_state.v1.ReadSideHandlerService",
	HandlerType: (*ReadSideHandlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleReadSide",
			Handler:    _ReadSideHandlerService_HandleReadSide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chief_of_state/v1/readside.proto",
}
