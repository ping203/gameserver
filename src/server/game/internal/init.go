package internal

import (
	"server/manager"

	"github.com/name5566/leaf/chanrpc"
)

var serverMgr *manager.ServerManager
var userMgr *userManager
var dbMgr *manager.DbManager

func Init(servers map[manager.ServerType]*chanrpc.Server) {

	serverMgr = &manager.ServerManager{}
	serverMgr.Init(servers)

	userMgr = &userManager{}
	userMgr.init()

	dbMgr = &manager.DbManager{}
	dbMgr.Init()
}
