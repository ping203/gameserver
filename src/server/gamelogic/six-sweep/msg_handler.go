package sixsweep

import (
	"reflect"

	"server/gamelogic"

	"github.com/golang/protobuf/proto"
)

type handler func(*gameSixSweep, gamelogic.User, proto.Message)

var gameMsgHandler = make(map[reflect.Type]handler)

func init() {

}

func register(message proto.Message, f handler) {
	typ := reflect.TypeOf(message)
	_, exist := gameMsgHandler[typ]
	if exist {
		panic("")
	}
}
