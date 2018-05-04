package internal

import (
	"server/gameproto/cmsg"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("checkAccount", checkAccount)
}

func checkAccount(args []interface{}) {
	resp := &cmsg.CRespAuth{}
	agent := args[1].(gate.Agent)
	if args[0] != nil {
		err := args[0].(error)
		if err != nil {
			resp.ErrCode = 1
			agent.WriteMsg(resp)
			return
		}
	}

	sessionMgr.
		agent.WriteMsg(resp)
}
