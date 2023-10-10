package server_packet

type RequestPacket struct {
	Data interface{} `json:"data"`
}

type ResponsePacket struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type ReqPacketTest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
type ResPacketTest struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Seconds string `json:"seconds"`
}
