package poke

import (
	"time"

	"server/gamelogic/fsm"
	"server/util"

	"github.com/golang/protobuf/proto"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

const chooseTime = time.Second * 20
const playTime = time.Second * 2
const clearTime = time.Second * 30

const (
	eventInit         = "event_init"
	eventStart        = "event_start"
	eventChoose       = "event_choose"
	eventPlayerAction = "event_player_action"
	eventGameOver     = "event_game_over"
)

const (
	stateInit         = "state_init"
	stateStart        = "state_start"
	stateChoose       = "state_choose"
	statePlayerAction = "state_player_action"
	stateGameOver     = "state_game_over"
)

func newGameFsm(g *GamePoke) *fsm.FSM {
	states := fsm.States{
		stateInit:         newStateInit(g),
		stateStart:        newStateStart(g),
		stateChoose:       newStateChoose(g),
		statePlayerAction: newStatePlayerAction(g),
		stateGameOver:     newStateGameOver(g),
	}
	f := fsm.NewFSM(stateInit, states, nil)
	return f
}

func newStateInit(g *GamePoke) fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventStart: stateStart,
		},
		OnEnter: func(e *fsm.Event) {

		},
		OnLeave: func(e *fsm.Event) {
		},
		InternalEvents: fsm.Callbacks{
			"start": func(e *fsm.Event) {
				g.fsm.Event(eventStart)
			},
		},
	}
}

func newStateStart(g *GamePoke) fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventChoose: stateChoose,
		},
		OnEnter: func(e *fsm.Event) {
			g.Start()
			generals := make([]*gamedef.GameGeneral, 0, len(g.players))
			for _, v := range g.players {
				g := v.Player.GameGeneral.getStatus(true)
				generals = append(generals, g)
			}
			g.notifyMessage(&cmsg.CNotifyGameStart{
				Generals: generals,
			})

			g.AfterPost(1, func() {
				g.fsm.Event(eventChoose)
			})
		},
		OnLeave: func(e *fsm.Event) {
		},
	}
}

func newStateChoose(g *GamePoke) fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventPlayerAction: statePlayerAction,
		},
		OnEnter: func(e *fsm.Event) {
			g.notifyGameStage(gamedef.GameStageTyp_GSTChoose, chooseTime)
			f := func() {
				for _, v := range g.players {
					if v.choose == nil {
						skillNum := len(v.GameGeneral.Skills)
						rand := util.RandNum(int32(skillNum))
						v.choose = &cmsg.CReqUseSkill{
							SkillID: v.GameGeneral.Skills[rand],
						}
					}
				}
				g.fsm.Event(eventPlayerAction)
			}
			g.AfterPost(chooseTime, f)
		},
		OnLeave: func(e *fsm.Event) {
		},
		InternalEvents: fsm.Callbacks{
			"choose": func(e *fsm.Event) {
				choose := e.Args[0].(proto.Message)
				p := e.Args[1].(*Player)

				p.choose = choose

				// 是否所有都做出选择
				for _, v := range g.players {
					if v.choose == nil {
						return
					}
				}
				// 取消超时处理
				g.stop()
				g.AfterPost(1, func() {
					g.fsm.Event(eventPlayerAction)
				})
			},
		},
	}
}

func newStatePlayerAction(g *GamePoke) fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventChoose:   stateChoose,
			eventGameOver: stateGameOver,
		},
		OnEnter: func(e *fsm.Event) {
			g.notifyGameStage(gamedef.GameStageTyp_GSTAction, playTime)
			// 伤害操作
			players := make(Players, 0, len(g.players))
			for _, v := range g.players {
				players = append(players, v)
			}
			g.sortPlayer(players)

			for _, v := range players {
				if !g.isStart() {
					v.chooseRoute(v.choose)
					if g.fsm.Current() != statePlayerAction {
						return
					}
					v.choose = nil
				}
			}

			if !g.isStart() {
				g.AfterPost(playTime, func() {
					g.fsm.Event(eventChoose)
				})
			}
		},
		OnLeave: func(e *fsm.Event) {
		},
		InternalEvents: fsm.Callbacks{
			"died": func(e *fsm.Event) {
				g.AfterPost(1, func() {
					g.fsm.Event(eventGameOver)
				})
			},
		},
	}
}

func newStateGameOver(g *GamePoke) fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventInit: stateInit,
		},
		OnEnter: func(e *fsm.Event) {
			for _, v := range g.players {
				exp := v.getOpponent().getExp(v.ID() == g.winner)
				result := &cmsg.CNotifyGameResult{
					Winner: g.winner,
					Exp:    exp,
				}
				v.SendMsg(result)
				v.AddExp(v.GameGeneral.PkID, exp)
			}
			g.clearUsers()
		},
		OnLeave: func(e *fsm.Event) {
		},
	}
}
