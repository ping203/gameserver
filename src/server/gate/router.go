package gate

import (
	"fmt"

	"server/gameproto/cmsg"
	"server/login"
	"server/msg"
)

// 初始化路由
func init() {
	bindClientMessage()
	sendClientMessage()
}

func bindClientMessage() {
	msg.Router(&cmsg.CReqAuth{}, login.ChanRPC)
	fmt.Println(&cmsg.CReqAuth{})
}

func sendClientMessage() {
	msg.Processor.Register(&cmsg.CRespAuth{})
}
