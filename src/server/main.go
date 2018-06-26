package main

import (
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
	"server/manager"

	"github.com/name5566/leaf"
	"github.com/name5566/leaf/chanrpc"
	lconf "github.com/name5566/leaf/conf"
	"github.com/sirupsen/logrus"
)

var servers = map[manager.ServerType]*chanrpc.Server{
	manager.GateServer:  gate.ChanRPC,
	manager.LoginServer: login.ChanRPC,
	manager.GameServer:  game.ChanRPC,
}

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	lconf.ConfigPath = conf.Server.ConfigFile

	logrus.SetLevel(logrus.Level(logrus.DebugLevel))

	text := new(logrus.TextFormatter)
	text.FullTimestamp = true
	logrus.SetFormatter(text)

	// 注册服务
	leaf.RegisterService(servers,
		game.Module,
		gate.Module,
		login.Module,
	)

	// 开始运行
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
