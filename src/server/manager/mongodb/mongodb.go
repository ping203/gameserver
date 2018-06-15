package mongodb

import (
	"server/gameproto/gamedef"
	"server/logs"

	"gopkg.in/mgo.v2-unstable"
)

type MgoClient struct {
	*mgo.Session
	dataBase string
}

type ModelUser struct {
	ID   uint64 `bson:"_id"`
	User *gamedef.User
}

// Init 初始化
func (m *MgoClient) Init(url string, dataBase string) {
	session, err := mgo.Dial(url)
	if err != nil {
		panic("dial mongodb err")
	}
	m.Session = session
	m.dataBase = dataBase
}

// Close 关闭连接
func (m *MgoClient) Close() {
	m.Session.Close()
}

// Find... 查找一条数据
func (m *MgoClient) Find(table string, id interface{}, result interface{}) (interface{}, error) {
	session := m.Session.Clone()
	defer session.Close()

	collection := session.DB(m.dataBase).C(table)
	err := collection.FindId(id).One(result)

	if err != nil {
		logs.Error("mongo_base method:Get " + err.Error())
		return nil, err
	}
	return result, nil
}

// Insert... 插入一条数据
func (m *MgoClient) Insert(table string, msg interface{}) error {
	session := m.Session.Clone()
	defer session.Close()

	collection := session.DB(m.dataBase).C(table)
	err := collection.Insert(msg)

	return err
}

// Update... 更新一条数据
func (m *MgoClient) Update(table string, id interface{}, msg interface{}) error {
	session := m.Session.Clone()
	defer session.Close()

	collection := session.DB(m.dataBase).C(table)
	err := collection.UpdateId(id, msg)

	return err
}
