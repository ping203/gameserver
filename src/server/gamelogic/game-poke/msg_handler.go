package poke

import (
	"fmt"
	"reflect"

	"server/gamelogic"
	"server/gameproto/cmsg"

	"github.com/golang/protobuf/proto"
)

type handler func(*GamePoke, gamelogic.User, proto.Message)

var gameMsgHandler = make(map[reflect.Type]handler)

func init() {
	register(useSkill)
}

func register(h interface{}) {
	v := reflect.ValueOf(h)

	msg := reflect.New(v.Type().In(2)).Elem().Interface().(proto.Message)

	typ := reflect.TypeOf(msg)
	_, exist := gameMsgHandler[typ]
	if exist {
		panic(fmt.Sprintf("message %v already register", msg))
	}

	gameMsgHandler[typ] = func(g *GamePoke, u gamelogic.User, msg proto.Message) {
		v.Call([]reflect.Value{reflect.ValueOf(g), reflect.ValueOf(u), reflect.ValueOf(msg)})
	}
}

func useSkill(g *GamePoke, u gamelogic.User, msg *cmsg.CReqUseSkill) error {
	player := g.findPlayByUserID(u.ID())

	err := player.useSkill(msg.SkillID)
	if err != nil {
		return err
	}
	return nil
}
