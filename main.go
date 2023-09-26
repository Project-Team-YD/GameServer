package main

import (
	"context"
	"net"
	global_grpc "project_yd/grpc"

	"google.golang.org/grpc"
	//itemcreateevent "github.com/heroiclabs/nakama/v3/nurhyme_common/ItemCreateEvent"
)

type server struct {
	global_grpc.UnimplementedGlobalGRpcServiceServer
}

func (s *server) SayHello(ctx context.Context, in *global_grpc.GlobalGrpcRequest) (*global_grpc.GlobalGrpcResponse, error) {
	println("Request!! Rpc Key::", in.RpcKey, "/ Message::", in.Message)
	result := &global_grpc.GlobalGrpcResponse{}

	result.Message = "Test Response Success"
	return result, nil
}

func main() {
	lis, err := net.Listen("tcp", ":19001")
	if err != nil {
		println(err)
	}

	grpcServer := grpc.NewServer()
	global_grpc.RegisterGlobalGRpcServiceServer(grpcServer, &server{})

	println("server listening at ", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		println("grpcServer Serv Err!!::", err)
	}
}
