package internal

import (
	"server/gameproto/cmsg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handler() {
	skeleton.RegisterHandler(onRespAuth)
}

func (p *Client) req() {
	p.reqAuth()
}

func (p *Client) reqAuth() {
	p.WriteMsg(&cmsg.CReqAuth{
		Account:  "xx",
		Password: "xx",
	})
}

func onRespAuth(msg *cmsg.CRespAuth, agent gate.Agent) {
	log.Debug("%v", msg)
}
