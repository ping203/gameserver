package internal

import (
	"server/manager"

	"github.com/name5566/leaf/chanrpc"
)

var serverMgr *manager.ServerManager

func Init(servers map[manager.ServerType]*chanrpc.Server) {

	serverMgr = &manager.ServerManager{}
	serverMgr.Init(servers)
}
