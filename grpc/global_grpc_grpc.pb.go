// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: protos/global_grpc.proto

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
	GlobalGRpcService_GlobalGRpc_FullMethodName       = "/main_grpc.GlobalGRpcService/GlobalGRpc"
	GlobalGRpcService_GlobalGrpcStream_FullMethodName = "/main_grpc.GlobalGRpcService/GlobalGrpcStream"
)

// GlobalGRpcServiceClient is the client API for GlobalGRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalGRpcServiceClient interface {
	GlobalGRpc(ctx context.Context, in *GlobalGrpcRequest, opts ...grpc.CallOption) (*GlobalGrpcResponse, error)
	GlobalGrpcStream(ctx context.Context, in *GlobalGrpcRequest, opts ...grpc.CallOption) (*GlobalGrpcResponse, error)
}

type globalGRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalGRpcServiceClient(cc grpc.ClientConnInterface) GlobalGRpcServiceClient {
	return &globalGRpcServiceClient{cc}
}

func (c *globalGRpcServiceClient) GlobalGRpc(ctx context.Context, in *GlobalGrpcRequest, opts ...grpc.CallOption) (*GlobalGrpcResponse, error) {
	out := new(GlobalGrpcResponse)
	err := c.cc.Invoke(ctx, GlobalGRpcService_GlobalGRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalGRpcServiceClient) GlobalGrpcStream(ctx context.Context, in *GlobalGrpcRequest, opts ...grpc.CallOption) (*GlobalGrpcResponse, error) {
	out := new(GlobalGrpcResponse)
	err := c.cc.Invoke(ctx, GlobalGRpcService_GlobalGrpcStream_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlobalGRpcServiceServer is the server API for GlobalGRpcService service.
// All implementations must embed UnimplementedGlobalGRpcServiceServer
// for forward compatibility
type GlobalGRpcServiceServer interface {
	GlobalGRpc(context.Context, *GlobalGrpcRequest) (*GlobalGrpcResponse, error)
	GlobalGrpcStream(context.Context, *GlobalGrpcRequest) (*GlobalGrpcResponse, error)
	mustEmbedUnimplementedGlobalGRpcServiceServer()
}

// UnimplementedGlobalGRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalGRpcServiceServer struct {
}

func (UnimplementedGlobalGRpcServiceServer) GlobalGRpc(context.Context, *GlobalGrpcRequest) (*GlobalGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalGRpc not implemented")
}
func (UnimplementedGlobalGRpcServiceServer) GlobalGrpcStream(context.Context, *GlobalGrpcRequest) (*GlobalGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalGrpcStream not implemented")
}
func (UnimplementedGlobalGRpcServiceServer) mustEmbedUnimplementedGlobalGRpcServiceServer() {}

// UnsafeGlobalGRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalGRpcServiceServer will
// result in compilation errors.
type UnsafeGlobalGRpcServiceServer interface {
	mustEmbedUnimplementedGlobalGRpcServiceServer()
}

func RegisterGlobalGRpcServiceServer(s grpc.ServiceRegistrar, srv GlobalGRpcServiceServer) {
	s.RegisterService(&GlobalGRpcService_ServiceDesc, srv)
}

func _GlobalGRpcService_GlobalGRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalGRpcServiceServer).GlobalGRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GlobalGRpcService_GlobalGRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalGRpcServiceServer).GlobalGRpc(ctx, req.(*GlobalGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalGRpcService_GlobalGrpcStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalGRpcServiceServer).GlobalGrpcStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GlobalGRpcService_GlobalGrpcStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalGRpcServiceServer).GlobalGrpcStream(ctx, req.(*GlobalGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GlobalGRpcService_ServiceDesc is the grpc.ServiceDesc for GlobalGRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GlobalGRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main_grpc.GlobalGRpcService",
	HandlerType: (*GlobalGRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GlobalGRpc",
			Handler:    _GlobalGRpcService_GlobalGRpc_Handler,
		},
		{
			MethodName: "GlobalGrpcStream",
			Handler:    _GlobalGRpcService_GlobalGrpcStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/global_grpc.proto",
}
