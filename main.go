package main

import (
	"encoding/json"
	server "project_yd/server"
	packet "project_yd/server/server_packet"
	util "project_yd/server/util"
	//itemcreateevent "github.com/heroiclabs/nakama/v3/nurhyme_common/ItemCreateEvent"
)

func RpcKeyHandlerFunc (payload string) string {
	requestPacket := packet.ReqPacketTest{}
	err := json.Unmarshal([]byte(payload), &requestPacket)
	if err != nil {
		return util.ResponseErrorMessage(400, err.Error())
	}
	responsePacket := packet.ResPacketTest{}
	responsePacket.Code = 200
	responsePacket.Message = "Sueccess"
	responsePacket.Seconds = "씨발"
	
	return util.ResponseMessage(responsePacket)
}


func RegistRpc() {
	server.RegistRpc("rpcTest", RpcKeyHandlerFunc)
}

func main() {
	//-- start grpc server
	server.StartGrpcServer()
	RegistRpc()

	select{}
}
