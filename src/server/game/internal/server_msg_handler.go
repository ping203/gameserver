package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/emsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/smsg"
)

func init() {
	skeleton.RegisterHandler(onGtGsReqLogin)
	skeleton.RegisterHandler(onGtGsReqLogout)
}

func onGtGsReqLogin(req *smsg.GtGsReqLogin, agent gate.Agent) {
	cbk := func(u *user, err error) {
		resp := &smsg.GtGsRespLogin{
			SeqID: req.SeqID,
		}
		if err != nil {
			resp.ErrCode = uint32(emsg.BizErr_BE_LoadUserData)
			resp.ErrMsg = emsg.BizErr_BE_LoadUserData.String()
			u.Send2Gate(resp)
			return
		}

		u.login()
		resp.Account = u.account
		resp.User = u.info.User
		u.Send2Gate(resp)
	}
	userMgr.onUserEnter(req.UserID, req.Account, req.Extra, agent, cbk)
}

func onGtGsReqLogout(req *smsg.GtGsReqLogout, agent gate.Agent) {
	userID, ok := agent.UserData().(uint64)
	if !ok {
		serverMgr.Send2Gate(&smsg.GtGsRespLogout{
			IsClose: req.IsClose,
		}, agent)
	}
	if user, exist := userMgr.findUser(userID); exist {
		userMgr.onUserRemove(user.info.User.UserID)
		serverMgr.Send2Gate(&smsg.GtGsRespLogout{
			IsClose: req.IsClose,
			UserID:  userID,
		}, agent)
	}
}
