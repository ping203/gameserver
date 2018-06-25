package gamelogic

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type User interface {
	// SendMsg向玩家发送消息
	SendMsg(proto.Message)
	// ID 获取Uid
	ID() uint64
	// IsRobot 是否机器人
	IsRobot() bool
	GetData() *gamedef.User
	UseItem(uint32) bool
	GetGeneral() *gamedef.General
}

// Service服务
type Service interface {
	Post(func())
	AfterPost(time.Duration, func())
}

type Game interface {
	// MsgRoute 消息处理
	MsgRoute(proto.Message)

	UserJoin(User) error
	UserQuit(User) error

	IsEmpty() bool

	GameStart() error
	ReqGameRecord(User)
	ReportGameStart()
	ReportGameEnd()
	ReportGameClear()
	SendMsgBatch(msg proto.Message, users []User)
	GetGameID() uint32
}
