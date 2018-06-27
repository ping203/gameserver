package internal

import (
	"server/base"
	"server/manager"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {
	userMgr.close()
}

func (m *Module) RegisterService(servers map[manager.ServerType]*chanrpc.Server) {
	Init(servers)
}
