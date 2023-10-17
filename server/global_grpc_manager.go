package server

import (
	"context"
	"io"
	"net"
	global_grpc "project_yd/grpc"
	"sync"

	"google.golang.org/grpc"
)

var GlobalGrpcEvent *sync.Map

// 함수 시그니처: 클라이언트로부터 받은 key 값에 따라 호출할 함수들
type RpcKeyHandlerFunc func(payload string) string

// -- rpc로 쓰일 function 등록
func RegistRpc(rpcKey string, function RpcKeyHandlerFunc) {
	_, ok := GlobalGrpcEvent.Load(rpcKey)
	if ok {
		println("RPC KEY : ", rpcKey, " is Duplicate!")
		return
	}
	GlobalGrpcEvent.Store(rpcKey, function)
	println("RPC KEY : ", rpcKey, " is Regist Success!")
}

// -- 클라이언트로부터 호출된 rpc명으로 호출할 function을 찾고 payload값을 넘겨주어 결과값을 받는다.
func LoadRpc(rpcKey string, payload string) string {
	var result string
	rpcFunc, ok := GlobalGrpcEvent.Load(rpcKey)
	if !ok {
		result = "Not Found Rpc Key :" + rpcKey
		println(result)
		return result
	} else {
		if function, ok := rpcFunc.(RpcKeyHandlerFunc); ok {
			return function(payload)
		} else {
			result = "Result is not RpcKeyHandlerFunc Type"
			return result
		}
	}
}

type GrpcServer struct {
	global_grpc.UnimplementedGlobalGRpcServiceServer
}

// -- 클라이언트로부터 호출이 들어오면 rpcKey를 기준으로 등록된 function을 호출할것을 찾는다.
func (server *GrpcServer) GlobalGRpc(ctx context.Context, request *global_grpc.GlobalGrpcRequest) (*global_grpc.GlobalGrpcResponse, error) {
	result := &global_grpc.GlobalGrpcResponse{}
	result.Message = LoadRpc(request.RpcKey, request.Message)

	return result, nil
}

func (server *GrpcServer) GlobalGRpcStream(stream global_grpc.GlobalGRpcService_GlobalGrpcStreamServer) error {
	for {
		// 클라이언트로부터 메시지를 받음
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// rpcKey와 message를 사용하여 결과를 생성
		result := LoadRpc(request.RpcKey, request.Message)

		// 결과를 클라이언트에게 보냄
		response := &global_grpc.GlobalGrpcResponse{
			Message: result,
		}
		if err := stream.SendMsg(response); err != nil {
			return err
		}
	}
}

func GrpcServe(grpcServer *grpc.Server, listen net.Listener) {
	if err := grpcServer.Serve(listen); err != nil {
		println("ListenGrpc Error!!::", err)
	}
}

func StartGrpcServer() {
	println("Start GRPC Server")
	GlobalGrpcEvent = new(sync.Map)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", SERVER_PORT)
	if err != nil {
		println("ListenError!!", err.Error())
		return
	}

	global_grpc.RegisterGlobalGRpcServiceServer(grpcServer, &GrpcServer{})
	go GrpcServe(grpcServer, lis)

}
