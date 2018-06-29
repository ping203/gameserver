package poke

import (
	"fmt"
	"reflect"

	"server/gamelogic"
	"server/gameproto/cmsg"

	"github.com/golang/protobuf/proto"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/emsg"
)

type handler func(*GamePoke, gamelogic.User, proto.Message)

var gameMsgHandler = make(map[reflect.Type]handler)

func init() {
	register(chooseUseSkill)
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

func chooseUseSkill(g *GamePoke, u gamelogic.User, msg *cmsg.CReqUseSkill) {
	resp := &cmsg.CRespUseSkill{}
	if g.fsm.Current() != stateChoose {
		resp.ErrCode = uint32(emsg.BizErr_BE_NotInStage)
		resp.ErrMsg = emsg.BizErr_BE_NotInStage.String()
		u.SendMsg(resp)
		return
	}

	// 做出选择
	player := g.findPlayByUserID(u.ID())

	if !player.GameGeneral.checkSkill(msg.SkillID) {
		resp.ErrCode = uint32(emsg.BizErr_BE_HasNoSkill)
		resp.ErrMsg = emsg.BizErr_BE_HasNoSkill.String()
		u.SendMsg(resp)
		return
	}

	g.fsm.Event("choose", msg, player)
	resp.SkillID = msg.SkillID
	u.SendMsg(resp)
}
