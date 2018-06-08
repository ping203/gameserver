package gate

import (
	"server/gameproto/cmsg"
	"server/msg"
)

// 初始化路由
func init() {
	bindClientMessage()
	sendClientMessage()
}

func bindClientMessage() {
	msg.Router(&cmsg.CReqAuth{}, ChanRPC)
	msg.Router(&cmsg.CReqLogin{}, ChanRPC)
}

func sendClientMessage() {
	msg.Processor.Register(&cmsg.CRespAuth{})
	msg.Processor.Register(&cmsg.CRespLogin{})
}
