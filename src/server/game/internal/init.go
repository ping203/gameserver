package internal

import (
	"time"

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
	dbMgr.Init("127.0.0.1:27017", "game1")

	cfgMgr = &manager.ConfManager{}
	cfgMgr.Init(&gameconf.GameConfigPathNode{
		BaseConfigPath: lconf.ConfigPath,
	})

}

func AfterPost(d time.Duration, f func()) func() {
	timer := skeleton.AfterFunc(d, f)
	return timer.Stop
}
