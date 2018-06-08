package internal

import (
	"server/logs"

	"github.com/name5566/leaf/log"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
)

func (p *Client) handler() {
	skeleton.RegisterHandlerClient(p.onRespAuth)
	skeleton.RegisterHandlerClient(p.onRespLogin)
}

func (p *Client) req() {
	p.reqAuth()
}

func (p *Client) reqAuth() {
	logs.Debug("=========reqAuth=========")
	p.WriteMsg(&cmsg.CReqAuth{
		Account:  "xx",
		Password: "xx",
	})
}

func (p *Client) onRespAuth(msg *cmsg.CRespAuth) {
	log.Debug("%v", msg)
	if msg.ErrCode == 0 {
		p.reqLogin(msg.Sign, msg.UserID)
	}
}

func (p *Client) reqLogin(sign string, userID uint64) {
	logs.Debug("=========reqLogin=========")
	p.WriteMsg(&cmsg.CReqLogin{
		Sign:   sign,
		UserID: userID,
	})
}

func (p *Client) onRespLogin(msg *cmsg.CRespLogin) {
	log.Debug("%v", msg)
}
