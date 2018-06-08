package internal

import (
	"server/gameproto/emsg"
	"server/gameproto/smsg"

	"github.com/name5566/leaf/gate"
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
			u.send2Gate(resp)
		}

		u.login()
		resp.Account = u.account
		resp.User = u.info
		u.send2Gate(resp)
	}
	userMgr.onUserEnter(req.UserID, req.Account, req.Extra, agent, cbk)
}
