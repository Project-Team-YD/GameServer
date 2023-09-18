package main

import (
	"context"
	"net"
	proto_test "project_yd/protofiles"

	"google.golang.org/grpc"
	//itemcreateevent "github.com/heroiclabs/nakama/v3/nurhyme_common/ItemCreateEvent"
)

type server struct {
	proto_test.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto_test.HelloRequest) (*proto_test.HelloResponse, error) {
	println("Test :: Req :: ", in.GetName())
	result := &proto_test.HelloResponse{}

	result.Message = "Test Response Success"
	return result, nil
}

func main() {
	lis, err := net.Listen("tcp", ":19001")
	if err != nil {
		println(err)
	}

	grpcServer := grpc.NewServer()
	proto_test.RegisterGreeterServer(grpcServer, &server{})

	println("server listening at ", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		println("grpcServer Serv Err!!::", err)
	}
}
