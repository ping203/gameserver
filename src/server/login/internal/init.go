package internal

import (
	"server/login/internal/model"
	"server/manager"

	"github.com/name5566/leaf/chanrpc"
)

var (
	accountModel *model.AccountModel
	serverMgr    *manager.ServerManager
)

type db struct {
}

func (p *db) CheckAccount(string, string) (uint64, error) {
	return 1, nil
}

func Init(servers map[manager.ServerType]*chanrpc.Server) {
	accountModel = &model.AccountModel{}
	a := &db{}
	accountModel.Init(a, skeleton)

	serverMgr = &manager.ServerManager{}
	serverMgr.Init(servers)
}

func Post(f func()) {
	skeleton.Post(f)
}
