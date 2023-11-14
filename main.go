package main

import (
	"project_yd/game"
	server "project_yd/server"
	"sync"
	//itemcreateevent "github.com/heroiclabs/nakama/v3/nurhyme_common/ItemCreateEvent"
)

func RegistRpc() {
	game.RegistSessionRpc()
}

func main() {

	var waitGroup sync.WaitGroup

	server.StartDBConnection()
	server.RedisConnection()
	//-- GoRoutine Count +1
	waitGroup.Add(1)
	go func() {
		//-- GoRoutine Count -1
		defer waitGroup.Done()
		//-- Start Grpc server
		server.StartGrpcServer()
		//-- Rpc 등록
		RegistRpc()
		//-- Notification 서버 연결
		server.ConnectToNotificationServer()
	}()

	waitGroup.Wait()
	select {}
}
