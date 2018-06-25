package poke

import (
	"fmt"
	"reflect"

	"server/gamelogic"

	"github.com/golang/protobuf/proto"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
)

var gameChooseHandler = make(map[reflect.Type]handler)

func init() {
	chooseRegister(useSkill)
}

func chooseRegister(h interface{}) {
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

func useSkill(g *GamePoke, u gamelogic.User, msg *cmsg.CReqUseSkill) {
	if g.fsm.Current() != statePlayerAction {
		panic("err state")
		return
	}

	// 使用技能
	player := g.findPlayByUserID(u.ID())
	player.useSkill(msg.SkillID)
}
