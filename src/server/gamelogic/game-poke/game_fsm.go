package poke

import (
	"server/gamelogic/fsm"

	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

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
		eventInit:         newStateInit(g),
		eventStart:        newStateStart(g),
		eventChoose:       newStateChoose(g),
		eventPlayerAction: newStatePlayerAction(g),
		eventGameOver:     newStateGameOver(g),
	}
	f := fsm.NewFSM(stateStart, states, nil)
	return f
}

func newStateInit(g *GamePoke) fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventInit: stateInit,
		},
		OnEnter: func(e *fsm.Event) {
		},
		OnLeave: func(e *fsm.Event) {
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

			g.fsm.Event(eventChoose)
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

		},
		OnLeave: func(e *fsm.Event) {
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
			g.fsm.Event(eventChoose)
		},
		OnLeave: func(e *fsm.Event) {
		},
	}
}

func newStateGameOver(g *GamePoke) fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventStart: stateStart,
		},
		OnEnter: func(e *fsm.Event) {
		},
		OnLeave: func(e *fsm.Event) {
		},
	}
}
