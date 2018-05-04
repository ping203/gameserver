package internal

import (
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
)

func init() {
}

type messageHandler func(proto.Message, gate.Agent)

func handler(h interface{}) {
	v := reflect.ValueOf(h)

	if v.Type().NumIn() != 2 {
		panic("handler params num wrong")
	}
	msg := reflect.New(v.Type().In(0)).Elem().Interface().(proto.Message)
	f := func(args []interface{}) {
		// 收到的 Hello 消息
		m := args[0].(proto.Message)
		// 消息的发送者
		a := args[1].(gate.Agent)
		// 调用
		v.Call([]reflect.Value{reflect.ValueOf(m), reflect.ValueOf(a)})
	}
	skeleton.RegisterChanRPC(reflect.TypeOf(msg), f)
}
