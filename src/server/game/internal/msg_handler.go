package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
)

func init() {
	skeleton.RegisterHandler(onReqUserInit)
	skeleton.RegisterHandler(onReqNotifyUserData)
}

func onReqUserInit(req *cmsg.CReqUserInit, agent gate.Agent) {
	if user, exist := userMgr.findUser(agent.UserData().(uint64)); exist {
		user.onReqUserInit(req)
	}
}

func onReqNotifyUserData(req *cmsg.CReqNotifyUserData, agent gate.Agent) {
	if user, exist := userMgr.findUser(agent.UserData().(uint64)); exist {
		user.onReqNotifyUserData(req)
	}
}
