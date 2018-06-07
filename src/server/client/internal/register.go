package internal

import (
	"server/msg"

	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
)

// 初始化路由
func (p *Client) init() {
	p.bindClientMessage()
	p.sendClientMessage()

	handler()
}

func (p *Client) bindClientMessage() {
	msg.Router(&cmsg.CRespAuth{}, ChanRPC)
}

func (p *Client) sendClientMessage() {
	msg.Processor.Register(&cmsg.CReqAuth{})
}
