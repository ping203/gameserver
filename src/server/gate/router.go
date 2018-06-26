package gate

import (
	"server/msg"

	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/smsg"
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
	msg.Processor.Register(&smsg.GtLsRespAuth{})
	msg.Processor.Register(&smsg.GtGsRespLogin{})
}
