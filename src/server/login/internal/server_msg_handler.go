package internal

import (
	"server/gameproto/smsg"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterHandler(onGtLsReqAuth)
}

func onGtLsReqAuth(msg *smsg.GtLsReqAuth, agent gate.Agent) {
	f := func(userID uint64, err error) {
		resp := &smsg.GtLsRespAuth{}
		if err != nil {

		}
		resp.Account = msg.Account
		resp.UserID = userID
		serverMgr.Send2Gate(resp, agent)
	}
	accountModel.CheckAccountAsync(msg.Account, msg.Password, f)
}
