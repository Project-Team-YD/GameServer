package request_packet

type HeartBeat struct {
	HeartBeat string `json:"heartBeat"`
}

type DuplicateLoginFromNotificationServer struct {
	UUID string `json:"uuid"`
}
