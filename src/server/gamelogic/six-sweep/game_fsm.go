package sixsweep

import "server/gamelogic/fsm"

const (
	eventStart    = "event_start"
	eventPlay     = "event_play"
	eventGameOver = "event_game_over"
)

const (
	stateStart    = "state_start"
	statePlay     = "state_play"
	stateGameOver = "state_game_over"
)

func newGameSixSweepFsm() *fsm.FSM {
	states := fsm.States{
		eventStart:    newStateStart(),
		eventPlay:     newStatePlay(),
		eventGameOver: newStateGameOver(),
	}
	f := fsm.NewFSM(stateStart, states, nil)
	return f
}

func newStateStart() fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventPlay: statePlay,
		},
		OnEnter: func(e *fsm.Event) {
		},
		OnLeave: func(e *fsm.Event) {
		},
	}
}

func newStatePlay() fsm.State {
	return fsm.State{
		Transitions: fsm.Transitions{
			eventGameOver: stateGameOver,
		},
		OnEnter: func(e *fsm.Event) {
		},
		OnLeave: func(e *fsm.Event) {
		},
	}
}

func newStateGameOver() fsm.State {
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
