package sixsweep

import (
	"server/gamelogic"
	"server/gamelogic/fsm"
)

const side = 6

type gameSixSweep struct {
	Player []*Player
	gamelogic.Service

	fsm        *fsm.FSM
	chessboard []*chess
}

func NewGameSixSweep(svc gamelogic.Service) *gameSixSweep {
	g := &gameSixSweep{}
	g.fsm = newGameSixSweepFsm()
	g.Service = svc
	return g
}

func (p *gameSixSweep) Init(users []gamelogic.User) {
	p.Player = make([]*Player, 0, len(users))
	for _, v := range users {
		p.Player = append(p.Player, newPlayer(v))
	}
}
