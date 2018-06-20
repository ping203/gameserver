package gamelogic

import (
	"server/gameproto/gamedef"

	"github.com/golang/protobuf/proto"
)

type User interface {
	// SendMsg向玩家发送消息
	SendMsg(proto.Message)
	// ID 获取Uid
	ID() uint64
	// IsRobot 是否机器人
	IsRobot() bool
	GetData() *gamedef.User
}

// Player 玩家
type Player interface {
	User
	OnReconnect()
	OnDisconnect()
}

// Service服务
type Service interface {
	Post(func())
}

type Game interface {
	// MsgRoute 消息处理
	MsgRoute(proto.Message)
	GameStart([]User) error
	ReqGameRecord(User) proto.Message
	ReportGameStart()
	ReportGameEnd()
	ReportGameClear()
	SendMsgBatch(msg proto.Message, users []User)
}
