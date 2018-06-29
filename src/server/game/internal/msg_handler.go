package internal

import (
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
)

func init() {
	skeleton.RegisterClientHandler(onReqUserInit)
	skeleton.RegisterClientHandler(onReqNotifyUserData)
	skeleton.RegisterClientHandler(onReqStageFight)
	skeleton.RegisterClientHandler(onReqUseSkill)
}

func onReqUserInit(req *cmsg.CReqUserInit, userID uint64) {
	if user, exist := userMgr.findUser(userID); exist {
		user.onReqUserInit(req)
	}
}

func onReqNotifyUserData(req *cmsg.CReqNotifyUserData, userID uint64) {
	if user, exist := userMgr.findUser(userID); exist {
		user.onReqNotifyUserData(req)
	}
}

func onReqStageFight(req *cmsg.CReqStageFight, userID uint64) {
	if user, exist := userMgr.findUser(userID); exist {
		user.onReqStageFight(req)
	}
}

func onReqUseSkill(req *cmsg.CReqUseSkill, userID uint64) {
	if user, exist := userMgr.findUser(userID); exist {
		user.onReqUseSkill(req)
	}
}
