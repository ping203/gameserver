package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/smsg"
)

func init() {
	skeleton.RegisterHandler(onGtLsReqAuth)
}

func onGtLsReqAuth(msg *smsg.GtLsReqAuth, agent gate.Agent) {
	f := func(account *gamedef.Account, err error) {
		resp := &smsg.GtLsRespAuth{}
		if err != nil {
			resp.ErrCode = 1
			serverMgr.Send2Gate(resp, agent)
			return
		}
		resp.Account = msg.Account
		resp.UserID = account.UserID
		serverMgr.Send2Gate(resp, agent)
	}
	dbMgr.LoadAccountAsync(msg.Account, f)
}
