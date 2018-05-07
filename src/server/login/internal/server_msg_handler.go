package internal

import (
	"server/gameproto/smsg"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterHandler(onGtLsReqAuth)
	skeleton.RegisterHandler(onGtGsReqLogin)
}

func onGtLsReqAuth(msg *smsg.GtLsReqAuth, agent gate.Agent) {

}

func onGtGsReqLogin(msg *smsg.GtGsReqLogin, agent gate.Agent) {

}
