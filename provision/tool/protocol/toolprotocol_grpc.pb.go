// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: provision/tool/protocol/toolprotocol.proto

package protocol

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

// InvocationServiceClient is the client API for InvocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvocationServiceClient interface {
	Invoke(ctx context.Context, in *ToolRequest, opts ...grpc.CallOption) (*ToolResponse, error)
}

type invocationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvocationServiceClient(cc grpc.ClientConnInterface) InvocationServiceClient {
	return &invocationServiceClient{cc}
}

func (c *invocationServiceClient) Invoke(ctx context.Context, in *ToolRequest, opts ...grpc.CallOption) (*ToolResponse, error) {
	out := new(ToolResponse)
	err := c.cc.Invoke(ctx, "/foundation.provision.tool.protocol.InvocationService/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvocationServiceServer is the server API for InvocationService service.
// All implementations should embed UnimplementedInvocationServiceServer
// for forward compatibility
type InvocationServiceServer interface {
	Invoke(context.Context, *ToolRequest) (*ToolResponse, error)
}

// UnimplementedInvocationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInvocationServiceServer struct {
}

func (UnimplementedInvocationServiceServer) Invoke(context.Context, *ToolRequest) (*ToolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

// UnsafeInvocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvocationServiceServer will
// result in compilation errors.
type UnsafeInvocationServiceServer interface {
	mustEmbedUnimplementedInvocationServiceServer()
}

func RegisterInvocationServiceServer(s grpc.ServiceRegistrar, srv InvocationServiceServer) {
	s.RegisterService(&InvocationService_ServiceDesc, srv)
}

func _InvocationService_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvocationServiceServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foundation.provision.tool.protocol.InvocationService/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvocationServiceServer).Invoke(ctx, req.(*ToolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InvocationService_ServiceDesc is the grpc.ServiceDesc for InvocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "foundation.provision.tool.protocol.InvocationService",
	HandlerType: (*InvocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _InvocationService_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provision/tool/protocol/toolprotocol.proto",
}
