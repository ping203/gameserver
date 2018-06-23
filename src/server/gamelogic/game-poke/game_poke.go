package poke

import (
	"server/gamelogic"
	"server/gamelogic/fsm"
	"server/manager"
)

const side = 6

type GamePoke struct {
	players map[uint64]*Player
	gamelogic.Service

	fsm *fsm.FSM
	*manager.ConfManager
}

func NewGameSixSweep(svc gamelogic.Service, cfg *manager.ConfManager) *GamePoke {
	g := &GamePoke{}
	g.fsm = newGameSixSweepFsm()
	g.Service = svc
	g.ConfManager = cfg
	return g
}

func (p *GamePoke) Init(users []User) error {
	p.players = make(map[uint64]*Player, len(users))
	for _, v := range users {
		player, err := newPlayer(v, p)
		if err != nil {
			return err
		}
		p.players[v.ID()] = player
	}
	return nil
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
