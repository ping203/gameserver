package sixsweep

import "server/gamelogic"

const playerCount = 2

type Player struct {
	gamelogic.User
	chessboard [][]*chess
}

func newPlayer(user gamelogic.User) *Player {
	player := &Player{}
	player.User = user
	return player
}

func (p *Player) OnReconnect() {

}

func (p *Player) OnDisconnect() {

}
