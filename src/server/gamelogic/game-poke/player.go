package poke

import (
	"server/gamelogic"
)

const playerCount = 2

type Player struct {
	*GamePoke
	gamelogic.User

	GameGeneral
}

func newPlayer(user gamelogic.User, poke *GamePoke) (*Player, error) {
	player := &Player{}
	player.User = user
	player.GamePoke = poke
	err := player.initGeneral()
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (p *Player) initGeneral() error {
	general := p.GetGeneral()
	gg, err := newGameGeneral(general, p)
	if err != nil {
		return err
	}
	p.GameGeneral = *gg
	return nil
}

func (p *Player) useSkill(skill uint32) error {
	op := p.getOpponent()
	err := p.GameGeneral.useSkill(skill, &op.GameGeneral)
	return err
}

// 获取对手
func (p *Player) getOpponent() *Player {
	for userID, v := range p.GamePoke.players {
		if p.ID() != userID {
			return v
		}
	}
	panic("getOpponent")
}

func (p *Player) OnReconnect() {

}

func (p *Player) OnDisconnect() {

}
