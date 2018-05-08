package manager

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
)

type ServerType uint32

const (
	GateServer  ServerType = 1
	LoginServer ServerType = 2
	GameServer  ServerType = 3
	End         ServerType = 4
)

var ErrServer = errors.New("no server")

type ServerManager struct {
	server map[ServerType]*chanrpc.Server
}

// Init 初始化
func (p *ServerManager) Init(servers map[ServerType]*chanrpc.Server) {
	p.server = make(map[ServerType]*chanrpc.Server, End-1)
	if servers != nil {
		p.server = servers
	}
}

//RegisterServer 注册服务
func (p *ServerManager) RegisterServer(typ ServerType, rpc *chanrpc.Server) {
	p.server[typ] = rpc
}

//Send2Game 调用Game
func (p *ServerManager) Send2Game(msg proto.Message, agent gate.Agent) error {
	s, exist := p.server[GameServer]
	if !exist {
		return ErrServer
	}
	s.GoProto(msg, agent)
	return nil
}

//Send2Login 调用Login
func (p *ServerManager) Send2Login(msg proto.Message, agent gate.Agent) error {
	s, exist := p.server[LoginServer]
	if !exist {
		return ErrServer
	}
	s.GoProto(msg, agent)
	return nil
}

//Send2Gate 调用Gate
func (p *ServerManager) Send2Gate(msg proto.Message, agent gate.Agent) error {
	s, exist := p.server[GateServer]
	if !exist {
		return ErrServer
	}
	s.GoProto(msg, agent)
	return nil
}
