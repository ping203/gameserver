package internal

import (
	"fmt"
	"reflect"

	"server/util"

	"github.com/golang/protobuf/proto"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type handler func(*aiUser, proto.Message)

type aiManager struct {
	aiID2ai map[uint64]*aiUser

	aiHandler map[reflect.Type]handler
}

func (p *aiManager) init() {
	p.aiID2ai = make(map[uint64]*aiUser)
	p.aiHandler = make(map[reflect.Type]handler)
	p.register()
}

func (p *aiManager) setAiUser(ai *aiUser) {
	p.aiID2ai[ai.aiID] = ai
}

func (p *aiManager) getAiUser(id uint64) (*aiUser, bool) {
	ai, exist := p.aiID2ai[id]
	if !exist {
		return nil, false
	}
	return ai, true
}

func (p *aiManager) delAiUser(id uint64) {
	delete(p.aiID2ai, id)
}

func (p *aiManager) idMaker() uint64 {
	return util.GeneratePKID()
}

func (p *aiManager) newAiUser(generalID uint32, level uint32) *aiUser {
	ai := newAiUser(p.idMaker())
	p.setAiUser(ai)
	ai.newGeneral(generalID, level)
	return ai
}

func (p *aiManager) register() {
	p.aiRegister(p.gameStage)
}

func (p *aiManager) route(ai *aiUser, msg proto.Message) {
	typ := reflect.TypeOf(msg)

	handler, exist := p.aiHandler[typ]
	if !exist {
		return
	}

	handler(ai, msg)
}

func (p *aiManager) aiRegister(h interface{}) {
	v := reflect.ValueOf(h)

	msg := reflect.New(v.Type().In(1)).Elem().Interface().(proto.Message)

	typ := reflect.TypeOf(msg)
	_, exist := p.aiHandler[typ]
	if exist {
		panic(fmt.Sprintf("message %v already register", msg))
	}

	p.aiHandler[typ] = func(u *aiUser, msg proto.Message) {
		v.Call([]reflect.Value{reflect.ValueOf(u), reflect.ValueOf(msg)})
	}
}

func (p *aiManager) gameStage(ai *aiUser, msg *cmsg.CNotifyGameStage) {
	switch msg.Stage {
	case gamedef.GameStageTyp_GSTChoose:
		skeleton.Post(func() { ai.useSkill() })
	}
}
