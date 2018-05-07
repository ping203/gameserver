package internal

import (
	"server/gameproto/cmsg"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterHandler(onReqAuth)
}

func onReqAuth(msg *cmsg.CReqAuth, agent gate.Agent) {
	accountModel.CheckAccountAsync(msg.Account, msg.Password, agent)
}

func onReqLogin(msg *cmsg.CReqLogin, agent gate.Agent) {
	resp := &cmsg.CRespLogin{}
	agent.WriteMsg(resp)
}
