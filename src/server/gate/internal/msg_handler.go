package internal

import (
	"fmt"

	"server/gameproto/cmsg"
	"server/gameproto/smsg"
	"server/login"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterHandler(onReqAuth)
}

func onReqAuth(msg *cmsg.CReqAuth, agent gate.Agent) {
	sessionMgr.addUserOnAuth(agent)
	fmt.Println("22222222222222222")
	login.ChanRPC.GoProto(&smsg.GtLsReqAuth{}, agent)
}

func onReqLogin(msg *cmsg.CReqLogin, agent gate.Agent) {
	resp := &cmsg.CRespLogin{}
	agent.WriteMsg(resp)
}
