package internal

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
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

func (p *aiUser) newGeneral(generalID uint32, level uint32) error {
	conf, exist := cfgMgr.GetConfig().GetGeneralConfByGeneralID(generalID)
	if !exist {
		return fmt.Errorf("general.chooseGeneral GetGeneralConfByGeneralID %v", generalID)
	}
	p.general = &gamedef.General{
		GeneralID:  1,
		Individual: util.RandIndividual(),
		Effort:     &gamedef.Individual{},
		Skills:     conf.BaseSkills,
		Level:      level,
	}

	return nil
}

// SendMsg向玩家发送消息
func (p *aiUser) SendMsg(msg proto.Message) {
	logrus.WithFields(
		logrus.Fields{
			"aiID": p.aiID,
			"msg":  msg,
		},
	).Debug("ai get msg")
	aiMgr.route(p, msg)
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

func (p *aiUser) SetGameID(gameID uint32) {
	p.gameID = gameID
}

func (p *aiUser) AddExp(pkID uint64, exp int32) {

}

func (p *aiUser) useSkill() {
	g, exist := gameMgr.getGameByID(p.gameID)
	if !exist {
		return
	}

	if len(p.general.Skills) == 0 {
		return
	}

	rand := util.RandNum(int32(len(p.general.Skills)))
	skill := p.general.Skills[int(rand)]

	logrus.WithFields(
		logrus.Fields{
			"aiID":    p.aiID,
			"skillID": skill,
		},
	).Debug("ai use skill")

	g.MsgRoute(&cmsg.CReqUseSkill{
		SkillID: skill,
	}, p)
}

func (p *aiUser) AddGeneral(gameGeneral *gamedef.GameGeneral) {

}
