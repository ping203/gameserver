package internal

import (
	"github.com/name5566/leaf/module"

	"server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
	closeSig []chan bool
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	m.closeSig = make([]chan bool, 2)
	for k := range m.closeSig {
		m.closeSig[k] = make(chan bool)
	}
}

func (m *Module) Run(closeSig chan bool) {
	m.Close(closeSig)
	m.Skeleton.Run(m.closeSig[0])
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
