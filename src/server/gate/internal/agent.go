package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/smsg"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	sessionMgr.addSession(a)
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	sess, exist := sessionMgr.getSessionByAgent(a)
	if exist && sess.userID != 0 {
		serverMgr.Send2Game(&smsg.GtGsReqLogout{
			IsClose: true,
		}, a)
	}
	sessionMgr.removeSession(a)
}
