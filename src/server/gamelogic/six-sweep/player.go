package sixsweep

import "server/gamelogic"

type Player struct {
	gamelogic.User
}

func NewPlayer(user gamelogic.User) *Player {
	player := &Player{}
	player.User = user
	return player
}

func (p *Player) OnReconnect() {

}

func (p *Player) OnDisconnect() {

}
