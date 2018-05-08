package internal

import (
	"server/gameproto/cmsg"
	"server/gameproto/smsg"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterHandler(onReqAuth)
}

func onReqAuth(msg *cmsg.CReqAuth, agent gate.Agent) {
	resp := &cmsg.CRespAuth{}
	session, exist := sessionMgr.getSessionByAgent(agent)
	if !exist {
		resp.ErrCode = 1
		agent.WriteMsg(resp)
		return
	}

	if session.isAuthing() {
		resp.ErrCode = 2
		agent.WriteMsg(resp)
		return
	}

	sessionMgr.addUserOnAuth(agent)
	err := serverMgr.Send2Login(&smsg.GtLsReqAuth{
		Account:  msg.Account,
		Password: msg.Password,
	}, agent)
	if err != nil {
		resp.ErrCode = 100
		agent.WriteMsg(resp)
		return
	}
}

func onReqLogin(msg *cmsg.CReqLogin, agent gate.Agent) {
	resp := &cmsg.CRespLogin{}
	agent.WriteMsg(resp)
}
