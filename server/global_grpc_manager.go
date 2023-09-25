package server

import grpc "project_yd/grpc"

type server_manager struct {
	grpc.UnimplementedGlobalRpcServiceServer
}

//func (s *server)
