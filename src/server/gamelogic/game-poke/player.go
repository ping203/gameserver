package poke

import (
	"server/gamelogic"
	"server/gameproto/gamedef"
)

const playerCount = 2

type User interface {
	gamelogic.User
	useItem(uint32) bool
	getGeneral() *gamedef.General
}

type Player struct {
	*GamePoke
	User

	GameGeneral
}

func newPlayer(user User, poke *GamePoke) (*Player, error) {
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
	general := p.getGeneral()
	gg, err := newGameGeneral(general, p)
	if err != nil {
		return err
	}
	p.GameGeneral = gg
}

func (p *Player) OnReconnect() {

}

func (p *Player) OnDisconnect() {

}
