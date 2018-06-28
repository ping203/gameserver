package internal

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
	"github.com/wenxiu2199/gameserver/src/server/util"
)

type aiUser struct {
	aiID uint64

	info    *gamedef.UserData
	general *gamedef.General

	// 游戏数据
	gameID uint32
}

func newAiUser(id uint64) *aiUser {
	ai := &aiUser{
		aiID: id,
		info: &gamedef.UserData{
			User: &gamedef.User{
				UserID:   id,
				Nickname: "robot",
			},
		},
	}

	return ai
}

func (p *aiUser) newGeneral(generalID uint32) error {
	conf, exist := cfgMgr.GetConfig().GetGeneralConfByGeneralID(generalID)
	if !exist {
		return fmt.Errorf("general.chooseGeneral GetGeneralConfByGeneralID %v", generalID)
	}
	p.general = &gamedef.General{
		GeneralID:  1,
		Individual: util.RandIndividual(),
		Effort:     &gamedef.Individual{},
		Skills:     conf.BaseSkills,
		Level:      5,
	}

	return nil
}

// SendMsg向玩家发送消息
func (p *aiUser) SendMsg(proto.Message) {

}

// ID 获取Uid
func (p *aiUser) ID() uint64 {
	return p.aiID
}

// IsRobot 是否机器人
func (p *aiUser) IsRobot() bool {
	return true
}
func (p *aiUser) GetData() *gamedef.User {
	return p.info.User
}
func (p *aiUser) UseItem(uint32) bool {
	return true
}
func (p *aiUser) GetGeneral() *gamedef.General {
	return p.general
}
