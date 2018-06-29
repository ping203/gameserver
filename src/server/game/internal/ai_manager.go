package internal

import (
	"fmt"
	"reflect"

	"server/gamelogic"
	"server/util"

	"github.com/golang/protobuf/proto"
)

type handler func(gamelogic.Game, *aiUser, proto.Message)

type aiManager struct {
	aiID2ai map[uint64]*aiUser

	aiHandler map[reflect.Type]handler
}

func (p *aiManager) init() {
	p.aiID2ai = make(map[uint64]*aiUser)
	p.aiHandler = make(map[reflect.Type]handler)
}

func (p *aiManager) aiRegister(h interface{}) {
	v := reflect.ValueOf(h)

	msg := reflect.New(v.Type().In(2)).Elem().Interface().(proto.Message)

	typ := reflect.TypeOf(msg)
	_, exist := p.aiHandler[typ]
	if exist {
		panic(fmt.Sprintf("message %v already register", msg))
	}

	p.aiHandler[typ] = func(g gamelogic.Game, u *aiUser, msg proto.Message) {
		v.Call([]reflect.Value{reflect.ValueOf(g), reflect.ValueOf(u), reflect.ValueOf(msg)})
	}
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

func (p *aiManager) newAiUser(generalID uint32) *aiUser {
	ai := newAiUser(p.idMaker())
	p.setAiUser(ai)
	ai.newGeneral(generalID)
	return ai
}
