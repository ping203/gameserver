package mongodb

import (
	"server/util"

	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
	"gopkg.in/mgo.v2-unstable"
)

const tbl = "Account"

type MgoAccount struct {
	*MgoClient
}

type ModelAccount struct {
	ID      string `bson:"_id"`
	Account *gamedef.Account
}

func (m *MgoAccount) Init(client *MgoClient) {
	m.MgoClient = client
}

func (m *MgoAccount) Find(account string) (*gamedef.Account, error) {
	item := &ModelAccount{}
	err := m.MgoClient.Find(tbl, account, item)
	return item.Account, err
}

func (m *MgoAccount) FindOrCreate(account string) (*gamedef.Account, error) {
	item := &ModelAccount{}
	err := m.MgoClient.Find(tbl, account, item)
	if err == mgo.ErrNotFound {
		account := &gamedef.Account{
			Account: account,
			UserID:  util.GeneratePKID(),
		}
		err := m.Create(account)
		return account, err
	}
	return item.Account, err
}

func (m *MgoAccount) Create(account *gamedef.Account) error {
	item := &ModelAccount{
		ID:      account.Account,
		Account: account,
	}
	err := m.MgoClient.Insert(tbl, item)
	return err
}

func (m *MgoAccount) Update(account *gamedef.Account) error {
	item := &ModelAccount{
		ID:      account.Account,
		Account: account,
	}
	err := m.MgoClient.Update(tbl, account.Account, item)
	return err
}
