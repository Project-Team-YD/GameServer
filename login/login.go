package login

import (
	"encoding/json"
	server "project_yd/server"
	request "project_yd/server/server_packet/request_packet"
	response "project_yd/server/server_packet/response_packet"
	"project_yd/util"
)

func RegistLoginRpc() {
	server.RegistRpc("login", LoginRpc)
}

func LoginRpc(payload string) string {
	requestPacket := request.Login{}
	err := json.Unmarshal([]byte(payload), &requestPacket)
	if err != nil {
		return util.ResponseErrorMessage(400, err.Error())
	}
	//---여기작업해야함
	//db := server.DBManager.Database[util.LOGIN]
	//loginSql := `SELECT * `

	responsePacket := response.Login{}
	responsePacket.MessageCode = 200
	responsePacket.Message = "Sueccess"

	return util.ResponseMessage(responsePacket)
}
