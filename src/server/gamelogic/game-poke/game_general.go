package poke

import (
	"fmt"

	"server/gameproto/gameconf"
	"server/gameproto/gamedef"
)

type GameGeneral struct {
	*Player
	generalID uint32
	// 基础血攻防特攻特防速度
	baseHp        int32
	baseAttack    int32
	baseDefense   int32
	baseSpAttack  int32
	baseSpDefense int32
	baseSpeed     int32

	curHP  int32 // 当前血量
	buff   map[gameconf.SkillEffectTyp]*gamedef.Buff
	skills []uint32
}

func newGameGeneral(general *gamedef.General, player *Player) (*GameGeneral, error) {
	gg := &GameGeneral{}
	gg.generalID = general.GeneralID
	gg.Player = player
	gg.buff = make(map[gameconf.SkillEffectTyp]*gamedef.Buff)
	gg.calculateBase(general)
	gg.curHP = gg.baseHp
	gg.skills = general.Skills

	return gg, nil
}

func (p *GameGeneral) calculateBase(general *gamedef.General) error {
	conf, exist := p.getConfig().GetConfig().GetGeneralConfByGeneralID(general.GeneralID)
	if !exist {
		return fmt.Errorf("calculateBase: GetGeneralConfByGeneralID %v", general.GeneralID)
	}

	p.baseHp = (conf.Hp*2+general.Individual.Hp+general.Effort.Hp/4)*int32(general.Level)/100 + int32(general.Level) + 10
	p.baseAttack = (conf.Atk*2+general.Individual.Attack+general.Effort.Attack/4)*int32(general.Level)/100 + 5
	p.baseDefense = (conf.Def*2+general.Individual.Defense+general.Effort.Defense/4)*int32(general.Level)/100 + 5
	p.baseSpAttack = (conf.Satk*2+general.Individual.SpAttack+general.Effort.SpAttack/4)*int32(general.Level)/100 + 5
	p.baseSpDefense = (conf.Sdef*2+general.Individual.SpDefense+general.Effort.SpDefense/4)*int32(general.Level)/100 + 5
	p.baseSpeed = (conf.Spd*2+general.Individual.Speed+general.Effort.Speed/4)*int32(general.Level)/100 + 5

	return nil
}

func (p *GameGeneral) checkSkill(skillID uint32) bool {
	for _, v := range p.skills {
		if v == skillID {
			return true
		}
	}
	return false
}

func (p *GameGeneral) useSkill(skillID uint32) bool {
	for _, v := range p.skills {
		if v == skillID {
			return true
		}
	}
	return false
}
