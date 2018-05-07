package msg

import (
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
	Processor.SetRouter(msg, msgRouter)
}
