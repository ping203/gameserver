package internal

import (
	"fmt"

	"server/gameproto/smsg"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterHandler(onGtLsRespAuth)
	skeleton.RegisterHandler(onGtGsRespLogin)
}

func onGtLsRespAuth(msg *smsg.GtLsRespAuth, agent gate.Agent) {
	fmt.Println(msg)
}

func onGtGsRespLogin(msg *smsg.GtGsRespLogin, agent gate.Agent) {

}
