package internal

import (
	"server/login/internal/model"
)

var sessionMgr *sessionManager
var accountModel *model.AccountModel

type db struct {
}

func (p *db) CheckAccount(string, string) error {
	return nil
}

func init() {
	accountModel = &model.AccountModel{}
	a := &db{}
	accountModel.Init(a, skeleton)
}
