package manager

import (
	"sync"

	"server/manager/mongodb"

	"github.com/name5566/leaf/chanrpc"
	"github.com/sirupsen/logrus"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type DbManager struct {
	worker *chanrpc.Server

	mongodb.MgoUser
	mongodb.MgoAccount

	*sync.WaitGroup
}

func (s *DbManager) Init(url string, dataBase string, service *chanrpc.Server) {
	s.worker = service
	s.WaitGroup = &sync.WaitGroup{}

	mgoClient := &mongodb.MgoClient{}
	mgoClient.Init(url, dataBase)

	s.MgoUser = mongodb.MgoUser{}
	s.MgoUser.Init(mgoClient)

	s.MgoAccount = mongodb.MgoAccount{}
	s.MgoAccount.Init(mgoClient)
}

func (s *DbManager) Run(closeSig chan bool) {
}

func (s *DbManager) Close() {
	s.WaitGroup.Wait()
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
	s.WaitGroup.Add(1)
	go func() {
		err := s.MgoUser.Update(&cp)
		if err != nil {
			logrus.Error("FlushUserAsync %v", cp)
		}
		s.WaitGroup.Done()
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
