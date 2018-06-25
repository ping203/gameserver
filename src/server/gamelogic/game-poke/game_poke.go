package poke

import (
	"server/gamelogic"
	"server/gamelogic/fsm"
	"server/manager"

	"github.com/golang/protobuf/proto"
)

const playerNm = 2

type GamePoke struct {
	players map[uint64]*Player
	gamelogic.Service
	gameID uint32

	winner uint64

	fsm *fsm.FSM
	*manager.ConfManager

	// 回合
	round uint32
}

func NewGame(svc gamelogic.Service, cfg *manager.ConfManager, gameID uint32) *GamePoke {
	g := &GamePoke{}
	g.fsm = newGameFsm(g)
	g.Service = svc
	g.ConfManager = cfg
	g.gameID = gameID
	g.players = make(map[uint64]*Player, playerNm)
	return g
}

func (p *GamePoke) Start() {
	for _, v := range p.players {
		v.initGeneral()
	}
}

func (p *GamePoke) getConfig() *manager.ConfManager {
	return p.ConfManager
}

func (p *GamePoke) findPlayByUserID(userID uint64) *Player {
	player, exist := p.players[userID]
	if !exist {
		panic("findPlayByUserID")
	}
	return player
}

// MsgRoute 消息处理
func (p *GamePoke) MsgRoute(proto.Message) {

}

// GameStart 游戏开始
func (p *GamePoke) GameStart() error {
	return nil
}

// ReqGameRecord...
func (p *GamePoke) ReqGameRecord(gamelogic.User) {

}

// ReportGameStart...
func (p *GamePoke) ReportGameStart() {

}

// ReportGameEnd..
func (p *GamePoke) ReportGameEnd() {

}

// ReportGameClear...
func (p *GamePoke) ReportGameClear() {

}

// SendMsgBatch...
func (p *GamePoke) SendMsgBatch(msg proto.Message, users []gamelogic.User) {

}

// UserJoin 玩家加入
func (p *GamePoke) UserJoin(user gamelogic.User) error {
	player, err := newPlayer(user, p)
	if err != nil {
		return err
	}
	p.players[user.ID()] = player

	return nil
}

// UserQuit 玩家加入
func (p *GamePoke) UserQuit(user gamelogic.User) error {
	//player := p.findPlayByUserID(user.ID())
	// 结束游戏 or 托管
	//if p.fsm.Current() == stateStart {
	//	delete(p.players,user.ID())
	//} else if p.fsm.Current() == statePlay {
	//
	//} else if p.fsm.Current() == stateGameOver {
	//	delete(p.players,user.ID())
	//}

	delete(p.players, user.ID())
	return nil
}

// GetGameID...
func (p *GamePoke) GetGameID() uint32 {
	return p.gameID
}

// IsEmpty 玩家加入
func (p *GamePoke) IsEmpty() bool {
	if len(p.players) == 0 {
		return true
	}

	return false
}

func (p *GamePoke) notifyMessage(msg proto.Message) {
	for _, v := range p.players {
		v.SendMsg(msg)
	}
}
