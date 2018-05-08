package model

import (
	"github.com/name5566/leaf/module"
)

type dbs interface {
	CheckAccount(string, string) (uint64, error)
}

type token struct {
	sign   string
	expire int64
}

type AccountModel struct {
	dbs
	skeleton *module.Skeleton
}

func (p *AccountModel) Init(dbs dbs, skeleton *module.Skeleton) {
	p.dbs = dbs
	p.skeleton = skeleton
	// todo超时回收token
}

func (p *AccountModel) CheckAccountAsync(account string, psw string, f func(uint64, error)) {
	go func() {
		userID, err := p.dbs.CheckAccount(account, psw)
		f(userID, err)
	}()
}
