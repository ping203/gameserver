package poke

import (
	"fmt"
	"reflect"

	"server/gamelogic"

	"github.com/golang/protobuf/proto"
)

type handler func(*GamePoke, gamelogic.User, proto.Message)

var gameMsgHandler = make(map[reflect.Type]handler)

func init() {

}

func register(message proto.Message, f handler) {
	typ := reflect.TypeOf(message)
	_, exist := gameMsgHandler[typ]
	if exist {
		panic(fmt.Sprintf("message %v already register", message))
	}

	gameMsgHandler[typ] = f
}
