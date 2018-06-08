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
}

func (p *Client) sendClientMessage() {
	msg.Processor.Register(&cmsg.CReqAuth{})
	msg.Processor.Register(&cmsg.CReqLogin{})
}
