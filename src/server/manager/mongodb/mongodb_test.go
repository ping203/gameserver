package mongodb

import (
	"fmt"
	"testing"

	"server/gameproto/gamedef"
)

func TestMongodb(t *testing.T) {
	dbMgr := &MgoClient{}
	dbMgr.Init("127.0.0.1:27017", "game1")

	err := dbMgr.Insert("user", &ModelUser{
		ID: 4556,
		User: &gamedef.User{
			Nickname: "asd",
		},
	})
	if err != nil {
		t.Error(err)
	}

	res := &ModelUser{}
	_, err = dbMgr.Find("user", 4556, res)
	fmt.Println(res)
	if err != nil {
		t.Error(err)
	}

	err = dbMgr.Update("user", 4556, &ModelUser{
		User: &gamedef.User{},
	})
	if err != nil {
		t.Error(err)
	}

}
