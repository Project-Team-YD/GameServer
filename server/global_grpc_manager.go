package server

import (
	"context"
	"net"
	global_grpc "project_yd/grpc"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	global_grpc.UnimplementedGlobalGRpcServiceServer
}

type GrpcMessage struct {
	RequestMessage *global_grpc.GlobalGrpcRequest
}

func (server *GrpcServer) ListenGrpc(ctx context.Context, request *global_grpc.GlobalGrpcRequest) (*global_grpc.GlobalGrpcResponse, error) {
	message := GrpcMessage{}
	message.RequestMessage = request

	result := &global_grpc.GlobalGrpcResponse{}
	return result, nil
}

func GrpcServe(grpcServer *grpc.Server, listen net.Listener) {
	if err := grpcServer.Serve(listen); err != nil {
		println("ListenGrpc Error!!::", err)
	}
}

func StartGrpcServer() {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", SERVER_PORT)
	if err != nil {
		println(err)
	}

	global_grpc.RegisterGlobalGRpcServiceServer(grpcServer, &GrpcServer{})
	go GrpcServe(grpcServer, lis)
}
