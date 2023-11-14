package response_packet

type HeartBeat struct {
	Code      uint   `json:"code"`
	Message   string `json:"message"`
	HeartBeat string `json:"heartBeat"`
}
