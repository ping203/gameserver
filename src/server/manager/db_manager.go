package manager

import (
	"server/manager/mongodb"

	"github.com/name5566/leaf/chanrpc"
	"github.com/sirupsen/logrus"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type DbManager struct {
	worker *chanrpc.Server
	close  chan bool

	mongodb.MgoUser
	mongodb.MgoAccount
}

func (s *DbManager) Init(url string, dataBase string) {
	s.worker = chanrpc.NewServer(10000)
	s.close = make(chan bool)
	s.Run(s.close)

	mgoClient := &mongodb.MgoClient{}
	mgoClient.Init(url, dataBase)

	s.MgoUser = mongodb.MgoUser{}
	s.MgoUser.Init(mgoClient)

	s.MgoAccount = mongodb.MgoAccount{}
	s.MgoAccount.Init(mgoClient)
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

func (s *DbManager) LoadUserAsync(userID uint64, cbk func(*gamedef.UserData, error)) {
	go func() {
		user, err := s.MgoUser.FindOrCreate(userID)
		s.worker.Post(func([]interface{}) {
			// todo 加载数据
			cbk(user, err)
		})
	}()
}

func (s *DbManager) FlushUserAsync(user *gamedef.UserData) {
	// 拷贝一份数据
	cp := *user
	go func() {
		err := s.MgoUser.Update(&cp)
		if err != nil {
			logrus.Error("FlushUserAsync %v", cp)
		}
	}()
}

func (s *DbManager) LoadAccountAsync(account string, cbk func(*gamedef.Account, error)) {
	go func() {
		user, err := s.MgoAccount.FindOrCreate(account)
		s.worker.Post(func([]interface{}) {
			// todo 加载数据
			cbk(user, err)
		})
	}()
}
