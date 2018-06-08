package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type user struct {
	gate.Agent
	// 用户数据
	account string
	info    *gamedef.User

	connected bool
}

func (p *user) login() {
	p.connected = true
}

func (p *user) logout() {
	p.connected = false
}

func (p *user) isLogin() bool {
	return p.connected
}

// GetAccount 获取账号.
func (p *user) GetAccount() string {
	return p.account
}

func (p *user) sendMsg(message proto.Message) {
	serverMgr.SendMsg("Send2Client", message, p.Agent)
}

func (p *user) send2Gate(message proto.Message) {
	serverMgr.Send2Gate(message, p.Agent)
}
