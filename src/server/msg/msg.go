package msg

import (
	"fmt"

	"server/protobuf"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/chanrpc"
)

var Processor = protobuf.NewProcessor()

func init() {
}

// Router 注册和路由
func Router(msg proto.Message, msgRouter *chanrpc.Server) {
	Processor.Register(msg)
	fmt.Println(msg.String())
	Processor.SetRouter(msg, msgRouter)
}
