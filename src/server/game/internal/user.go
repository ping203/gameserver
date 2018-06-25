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

	general

	connected bool

	// 游戏数据
	gameID uint32
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

func (p *user) setGameID(gameID uint32) {
	p.gameID = gameID
}

func (p *user) clearGameID() {
	p.gameID = 0
}

// GetAccount 获取账号.
func (p *user) GetAccount() string {
	return p.account
}

// SendMsg向玩家发送消息
func (p *user) SendMsg(message proto.Message) {
	serverMgr.SendMsg("Send2Client", message, p.Agent)
}

func (p *user) Send2Gate(message proto.Message) {
	serverMgr.Send2Gate(message, p.Agent)
}

// ID 获取Uid
func (p *user) ID() uint64 {
	return p.info.UserID
}

func (p *user) IsRobot() bool {
	return false
}

func (p *user) GetData() *gamedef.User {
	return p.info
}

func (p *user) UseItem(uint32) bool {
	return true
}

func (p *user) GetGeneral() *gamedef.General {
	g, exist := p.getFightGeneral()
	// 获取不到上阵武将说明进入游戏前逻辑有错误
	if !exist {
		panic("no user general data")
	}
	return g
}
