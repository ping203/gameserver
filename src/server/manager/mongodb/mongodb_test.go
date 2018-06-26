package mongodb

import (
	"fmt"
	"testing"

	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

func TestMongodb(t *testing.T) {
	dbMgr := &MgoClient{}
	dbMgr.Init("127.0.0.1:27017", "game1")

	//mongodbUser(t, dbMgr)
	mongodbAccount(t, dbMgr)

}

func mongodbUser(t *testing.T, client *MgoClient) {
	mgoUser := &MgoUser{
		MgoClient: client,
	}
	err := mgoUser.Create(&gamedef.User{
		UserID:   12,
		Nickname: "asd",
	})
	if err != nil {
		t.Error(err)
	}
	err = mgoUser.Update(&gamedef.User{
		UserID:   12,
		Nickname: "asdwwwww",
	})
	if err != nil {
		t.Error(err)
	}
	u, err := mgoUser.FindUser(22)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}

func mongodbAccount(t *testing.T, client *MgoClient) {
	mgoAccount := &MgoAccount{
		MgoClient: client,
	}
	err := mgoAccount.Create(&gamedef.Account{
		UserID:  12,
		Account: "asd",
	})
	if err != nil {
		t.Error(err)
	}
	err = mgoAccount.Update(&gamedef.Account{
		UserID:  122,
		Account: "asd",
	})
	if err != nil {
		t.Error(err)
	}
	u, err := mgoAccount.FindAccount("asd")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}
