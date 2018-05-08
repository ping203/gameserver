package internal

import (
	"server/manager"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"

	"server/base"
	"server/conf"
	"server/msg"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*gate.Gate
	*module.Skeleton
	closeSig []chan bool
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    ChanRPC,
	}
	m.Skeleton = skeleton
	m.closeSig = make([]chan bool, 2)
	for k := range m.closeSig {
		m.closeSig[k] = make(chan bool)
	}
}

func (m *Module) RegisterService(servers map[manager.ServerType]*chanrpc.Server) {
	Init(servers)
}

func (m *Module) Run(closeSig chan bool) {
	m.Close(closeSig)
	go func() {
		m.Skeleton.Run(m.closeSig[0])
	}()
	m.Gate.Run(m.closeSig[1])
}

func (m *Module) Close(closeSig chan bool) {
	go func() {
		select {
		case <-closeSig:
			for _, v := range m.closeSig {
				v <- true
			}
			break
		}
	}()
}
