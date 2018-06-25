package poke

import (
	"reflect"

	"server/gamelogic"

	"github.com/golang/protobuf/proto"
)

const playerCount = 2

type Players []*Player

func (p Players) Less(i, j int) bool {
	return p[i].GameGeneral.getSpeed() > p[j].GameGeneral.getSpeed()
}

func (p Players) Len() int {
	return len(p)
}

func (p Players) Swap(i, j int) {
	t := p[i]
	p[i] = p[j]
	p[j] = t
}

type Player struct {
	*GamePoke
	gamelogic.User

	GameGeneral

	// 本回合是否操作过
	choose proto.Message
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

// chooseRoute
func (p *Player) chooseRoute(msg proto.Message) {
	typ := reflect.TypeOf(msg)

	handler, exist := gameChooseHandler[typ]
	if !exist {
		return
	}

	handler(p.GamePoke, p.User, msg)
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

func (p *Player) setChoose(msg proto.Message) {
	p.choose = msg
}

func (p *Player) getChoose() proto.Message {
	return p.choose
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
