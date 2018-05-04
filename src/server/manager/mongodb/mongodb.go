package mongodb

import (
	"mgo"

	"github.com/astaxie/beego/logs"
)

var (
	DBMgr      = new(DBManager)
	mgoSession *mgo.Session
	dataBase   = "game1"
	url        = "127.0.0.1:27017"
)

type DBManager struct {
	*mgo.Session
}

func init() {
	session, err := mgo.Dial(url)
	if err != nil {
		panic("dial mongodb err")
	}
	DBMgr.Session = session
}

// Find... 查找一条数据
func (m *DBManager) Find(table string, id string, result interface{}) (interface{}, error) {
	session := m.Session.Clone()
	defer session.Close()

	collection := session.DB(dataBase).C(table)
	err := collection.FindId(id).One(result)

	if err != nil {
		logs.Error("mongo_base method:Get " + err.Error())
		return nil, err
	}
	return result, nil
}

// Insert... 插入一条数据
func (m *DBManager) Insert(table string, id string, msg interface{}) error {
	session := m.Session.Clone()
	defer session.Close()

	collection := session.DB(dataBase).C(table)
	err := collection.Insert(msg)

	return err
}

// Update... 更新一条数据
func (m *DBManager) Update(table string, id string, msg interface{}) error {
	session := m.Session.Clone()
	defer session.Close()

	collection := session.DB(dataBase).C(table)
	err := collection.UpdateId(id, msg)

	return err
}
