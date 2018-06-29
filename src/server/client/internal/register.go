package internal

import (
	"server/msg"

	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
)

// 初始化路由
func (p *Client) init() {
	p.bindClientMessage()
	p.sendClientMessage()

	p.handler()
}

func (p *Client) bindClientMessage() {
	msg.Router(&cmsg.CRespAuth{}, ChanRPC)
	msg.Router(&cmsg.CRespLogin{}, ChanRPC)
	msg.Router(&cmsg.CRespUserInit{}, ChanRPC)
	msg.Router(&cmsg.CRespStageFight{}, ChanRPC)
	msg.Router(&cmsg.CNotifyGameResult{}, ChanRPC)
	msg.Router(&cmsg.CNotifyGameStage{}, ChanRPC)
	msg.Router(&cmsg.CNotifyGameStart{}, ChanRPC)
	msg.Router(&cmsg.CRespUseSkill{}, ChanRPC)
}

func (p *Client) sendClientMessage() {
	msg.Processor.Register(&cmsg.CReqAuth{})
	msg.Processor.Register(&cmsg.CReqLogin{})

	msg.Processor.Register(&cmsg.CReqUserInit{})
	msg.Processor.Register(&cmsg.CReqStageFight{})
	msg.Processor.Register(&cmsg.CReqUseSkill{})
}
