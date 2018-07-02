package internal

import (
	"time"

	"server/manager"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"github.com/sirupsen/logrus"
)

var sessionMgr *sessionManager
var serverMgr *manager.ServerManager
var respDispatcher *manager.RespDispatcher
var requester manager.Requester

const requesterTimeOut = time.Second * 5

func Init(servers map[manager.ServerType]*chanrpc.Server) {
	sessionMgr = &sessionManager{}
	sessionMgr.init()

	serverMgr = &manager.ServerManager{}
	serverMgr.Init(servers)

	respDispatcher = manager.NewRespDispatcher(skeleton)
	requester = manager.NewRequester(respDispatcher)

}

func Post(f func()) {
	skeleton.Post(f)
}

func writeMsg(agent gate.Agent, msg proto.Message) {
	logrus.Debug("send2client:", msg.String())
	agent.WriteMsg(msg)
}
