package poke

import (
	"reflect"
	"sort"
	"time"

	"server/gamelogic"
	"server/gamelogic/fsm"
	"server/gameproto/cmsg"
	"server/manager"

	"github.com/golang/protobuf/proto"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

const playerNm = 2

type GamePoke struct {
	gamelogic.Service
	gameID uint32
	fsm    *fsm.FSM
	*manager.ConfManager
	*gameTimer

	// 回合
	round   uint32
	players map[uint64]*Player
	winner  uint64

	cancel func()
}

func NewGame(svc gamelogic.Service, cfg *manager.ConfManager, gameID uint32) *GamePoke {
	g := &GamePoke{}
	g.fsm = newGameFsm(g)
	g.Service = svc
	g.ConfManager = cfg
	g.gameID = gameID
	g.players = make(map[uint64]*Player, playerNm)

	g.gameTimer = NewGameTimer(func(f func()) {
		g.Service.Post(f)
	})

	return g
}

func (p *GamePoke) Post(f func()) {
	p.Service.Post(f)
}

func (p *GamePoke) AfterPost(d time.Duration, f func()) {
	p.gameTimer.start(d, f)
}

func (p *GamePoke) Stop() {
	p.gameTimer.stop()
}

func (p *GamePoke) Start() {
	for _, v := range p.players {
		v.initGeneral()
	}
}

func (p *GamePoke) notifyGameStage(stage gamedef.GameStageTyp, s int32) {
	p.notifyMessage(&cmsg.CNotifyGameStage{
		Stage:    stage,
		LastTime: s,
	})

	logrus.Debug("进入阶段:%s", stage.String())
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
func (p *GamePoke) MsgRoute(msg proto.Message, user gamelogic.User) {
	typ := reflect.TypeOf(msg)

	handler, exist := gameMsgHandler[typ]
	if !exist {
		return
	}

	handler(p, user, msg)
}

// GameStart 游戏开始
func (p *GamePoke) GameStart() error {
	p.fsm.Event("start")
	return nil
}

func (p *GamePoke) sortPlayer(players Players) {
	sort.Sort(players)
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
	if len(p.players) >= playerCount {
		return errors.New("err player num")
	}
	_, exist := p.players[user.ID()]
	if exist {
		return errors.New("user already in room")
	}
	player, err := newPlayer(user, p)
	if err != nil {
		return err
	}

	p.players[user.ID()] = player

	return nil
}

// UserReady 玩家加入
func (p *GamePoke) UserReady(user gamelogic.User, ready bool) error {
	player := p.findPlayByUserID(user.ID())
	player.setReady(ready)
	if ready && p.allReady() {
		p.GameStart()
	}
	return nil
}

func (p *GamePoke) allReady() bool {
	count := 0
	for _, v := range p.players {
		if v.ready {
			count++
		}
	}
	return count == playerCount
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
