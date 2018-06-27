package gate

import (
	"server/game"
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
	// gate
	msg.Router(&cmsg.CReqAuth{}, ChanRPC)
	msg.Router(&cmsg.CReqLogin{}, ChanRPC)

	// game
	msg.Router(&cmsg.CReqUserInit{}, game.ChanRPC)
	msg.Router(&cmsg.CReqNotifyUserData{}, game.ChanRPC)
}

func sendClientMessage() {
	msg.Processor.Register(&cmsg.CRespAuth{})
	msg.Processor.Register(&cmsg.CRespLogin{})
	msg.Processor.Register(&smsg.GtLsRespAuth{})
	msg.Processor.Register(&smsg.GtGsRespLogin{})
	msg.Processor.Register(&cmsg.CRespUserInit{})
	msg.Processor.Register(&cmsg.CNotifyDataChange{})
	msg.Processor.Register(&cmsg.CRespNotifyUserData{})
}
