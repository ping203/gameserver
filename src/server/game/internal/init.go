package internal

import (
	"server/config/gameconf"
	"server/manager"

	"github.com/name5566/leaf/chanrpc"
	lconf "github.com/name5566/leaf/conf"
)

var serverMgr *manager.ServerManager
var userMgr *userManager
var dbMgr *manager.DbManager

var cfgMgr *manager.ConfManager

func Init(servers map[manager.ServerType]*chanrpc.Server) {

	serverMgr = &manager.ServerManager{}
	serverMgr.Init(servers)

	userMgr = &userManager{}
	userMgr.init()

	dbMgr = &manager.DbManager{}
	dbMgr.Init()

	cfgMgr = &manager.ConfManager{}
	cfgMgr.Init(&gameconf.GameConfigPathNode{
		BaseConfigPath: lconf.ConfigPath,
	})

}
