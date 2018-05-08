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
	sessionMgr.addUserOnAuth(agent)
	err := serverMgr.Send2Login(&smsg.GtLsReqAuth{
		Account:  msg.Account,
		Password: msg.Password,
	}, agent)
	if err != nil {

	}
}

func onReqLogin(msg *cmsg.CReqLogin, agent gate.Agent) {
	resp := &cmsg.CRespLogin{}
	agent.WriteMsg(resp)
}
