// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: global_grpc.proto

package grpc

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

const (
	GlobalRpcService_SayHello_FullMethodName = "/main_grpc.GlobalRpcService/SayHello"
)

// GlobalRpcServiceClient is the client API for GlobalRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalRpcServiceClient interface {
	SayHello(ctx context.Context, in *GlobalGrpcRequest, opts ...grpc.CallOption) (*GlobalGrpcResponse, error)
}

type globalRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalRpcServiceClient(cc grpc.ClientConnInterface) GlobalRpcServiceClient {
	return &globalRpcServiceClient{cc}
}

func (c *globalRpcServiceClient) SayHello(ctx context.Context, in *GlobalGrpcRequest, opts ...grpc.CallOption) (*GlobalGrpcResponse, error) {
	out := new(GlobalGrpcResponse)
	err := c.cc.Invoke(ctx, GlobalRpcService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlobalRpcServiceServer is the server API for GlobalRpcService service.
// All implementations must embed UnimplementedGlobalRpcServiceServer
// for forward compatibility
type GlobalRpcServiceServer interface {
	SayHello(context.Context, *GlobalGrpcRequest) (*GlobalGrpcResponse, error)
	mustEmbedUnimplementedGlobalRpcServiceServer()
}

// UnimplementedGlobalRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalRpcServiceServer struct {
}

func (UnimplementedGlobalRpcServiceServer) SayHello(context.Context, *GlobalGrpcRequest) (*GlobalGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGlobalRpcServiceServer) mustEmbedUnimplementedGlobalRpcServiceServer() {}

// UnsafeGlobalRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalRpcServiceServer will
// result in compilation errors.
type UnsafeGlobalRpcServiceServer interface {
	mustEmbedUnimplementedGlobalRpcServiceServer()
}

func RegisterGlobalRpcServiceServer(s grpc.ServiceRegistrar, srv GlobalRpcServiceServer) {
	s.RegisterService(&GlobalRpcService_ServiceDesc, srv)
}

func _GlobalRpcService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRpcServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GlobalRpcService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRpcServiceServer).SayHello(ctx, req.(*GlobalGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GlobalRpcService_ServiceDesc is the grpc.ServiceDesc for GlobalRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GlobalRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main_grpc.GlobalRpcService",
	HandlerType: (*GlobalRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GlobalRpcService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "global_grpc.proto",
}
