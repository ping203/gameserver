package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/emsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/smsg"
)

func init() {
	skeleton.RegisterHandler(onGtGsReqLogin)
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
