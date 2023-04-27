// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: broadcast.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BroadcastService_SingleBroadcast_FullMethodName = "/openPanel.BroadcastService/SingleBroadcast"
	BroadcastService_MultiBroadcast_FullMethodName  = "/openPanel.BroadcastService/MultiBroadcast"
)

// BroadcastServiceClient is the client API for BroadcastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadcastServiceClient interface {
	SingleBroadcast(ctx context.Context, in *SingleBroadcastRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MultiBroadcast(ctx context.Context, in *MultiBroadcastRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type broadcastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastServiceClient(cc grpc.ClientConnInterface) BroadcastServiceClient {
	return &broadcastServiceClient{cc}
}

func (c *broadcastServiceClient) SingleBroadcast(ctx context.Context, in *SingleBroadcastRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BroadcastService_SingleBroadcast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastServiceClient) MultiBroadcast(ctx context.Context, in *MultiBroadcastRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BroadcastService_MultiBroadcast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServiceServer is the server API for BroadcastService service.
// All implementations should embed UnimplementedBroadcastServiceServer
// for forward compatibility
type BroadcastServiceServer interface {
	SingleBroadcast(context.Context, *SingleBroadcastRequest) (*emptypb.Empty, error)
	MultiBroadcast(context.Context, *MultiBroadcastRequest) (*emptypb.Empty, error)
}

// UnimplementedBroadcastServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBroadcastServiceServer struct {
}

func (UnimplementedBroadcastServiceServer) SingleBroadcast(context.Context, *SingleBroadcastRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SingleBroadcast not implemented")
}
func (UnimplementedBroadcastServiceServer) MultiBroadcast(context.Context, *MultiBroadcastRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiBroadcast not implemented")
}

// UnsafeBroadcastServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadcastServiceServer will
// result in compilation errors.
type UnsafeBroadcastServiceServer interface {
	mustEmbedUnimplementedBroadcastServiceServer()
}

func RegisterBroadcastServiceServer(s grpc.ServiceRegistrar, srv BroadcastServiceServer) {
	s.RegisterService(&BroadcastService_ServiceDesc, srv)
}

func _BroadcastService_SingleBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServiceServer).SingleBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BroadcastService_SingleBroadcast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServiceServer).SingleBroadcast(ctx, req.(*SingleBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadcastService_MultiBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServiceServer).MultiBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BroadcastService_MultiBroadcast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServiceServer).MultiBroadcast(ctx, req.(*MultiBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BroadcastService_ServiceDesc is the grpc.ServiceDesc for BroadcastService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BroadcastService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openPanel.BroadcastService",
	HandlerType: (*BroadcastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SingleBroadcast",
			Handler:    _BroadcastService_SingleBroadcast_Handler,
		},
		{
			MethodName: "MultiBroadcast",
			Handler:    _BroadcastService_MultiBroadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "broadcast.proto",
}
