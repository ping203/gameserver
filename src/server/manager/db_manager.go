package manager

import (
	"server/manager/mongodb"

	"github.com/name5566/leaf/chanrpc"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type DbManager struct {
	worker *chanrpc.Server
	close  chan bool

	*mongodb.MgoClient
}

func (s *DbManager) Init() {
	s.worker = chanrpc.NewServer(10000)
	s.close = make(chan bool)
	s.Run(s.close)

	s.MgoClient = &mongodb.MgoClient{}
}

func (s *DbManager) Run(closeSig chan bool) {
	go func() {
		for {
			select {
			case <-closeSig:
				s.worker.Close()
				return
			case ci := <-s.worker.ChanCall:
				s.worker.Exec(ci)
			}
		}
	}()
}

func (s *DbManager) Close() {
	s.close <- true
}

func (s *DbManager) LoadUserAsync(userID uint64, cbk func(*gamedef.User, error)) {
	s.worker.Post(func([]interface{}) {
		// todo 加载数据
		cbk(&gamedef.User{}, nil)
	})
}
