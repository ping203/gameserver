package internal

import (
	"fmt"

	"server/gameproto/smsg"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterHandler(onGtLsReqAuth)
	skeleton.RegisterHandler(onGtGsReqLogin)
}

func onGtLsReqAuth(msg *smsg.GtLsReqAuth, agent gate.Agent) {
	f := func(userID uint64, err error) {
		resp := &smsg.GtLsRespAuth{}
		if err != nil {

		}
		resp.Account = msg.Account
		resp.UserID = userID
		fmt.Println(resp)
		serverMgr.Send2Gate(resp, agent)
	}
	accountModel.CheckAccountAsync(msg.Account, msg.Password, f)
}

func onGtGsReqLogin(msg *smsg.GtGsReqLogin, agent gate.Agent) {

}
