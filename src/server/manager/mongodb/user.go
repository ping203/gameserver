package mongodb

import (
	"fmt"

	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
	"gopkg.in/mgo.v2-unstable"
)

type MgoUser struct {
	*MgoClient
}

type ModelUser struct {
	ID   uint64 `bson:"_id"`
	User *gamedef.UserData
}

func (m *MgoUser) Init(client *MgoClient) {
	m.MgoClient = client
}

func (m *MgoUser) Find(id uint64) (*gamedef.UserData, error) {
	mod := id % 10
	item := &ModelUser{}
	err := m.MgoClient.Find(fmt.Sprintf("User_%v", mod), id, item)
	return item.User, err
}

func (m *MgoUser) FindOrCreate(id uint64) (*gamedef.UserData, error) {
	mod := id % 10
	item := &ModelUser{}
	err := m.MgoClient.Find(fmt.Sprintf("User_%v", mod), id, item)
	if err == mgo.ErrNotFound {
		user := &gamedef.UserData{
			User: &gamedef.User{
				UserID: id,
			},
		}
		err := m.Create(user)
		return user, err
	}
	return item.User, err
}

func (m *MgoUser) Create(user *gamedef.UserData) error {
	mod := user.User.UserID % 10
	item := &ModelUser{
		ID:   user.User.UserID,
		User: user,
	}
	err := m.MgoClient.Insert(fmt.Sprintf("User_%v", mod), item)
	return err
}

func (m *MgoUser) Update(user *gamedef.UserData) error {
	mod := user.User.UserID % 10
	item := &ModelUser{
		ID:   user.User.UserID,
		User: user,
	}
	err := m.MgoClient.Update(fmt.Sprintf("User_%v", mod), item.ID, item)
	return err
}
