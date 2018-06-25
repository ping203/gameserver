package gameconf

import "github.com/wenxiu2199/gameserver/src/server/gameproto/gameconf"

type skillConf struct {
	cfg *GameConfig
}

func (p *skillConf) init(cfg *GameConfig) error {
	p.cfg = cfg

	return nil
}

func (p *skillConf) GetSkillConfBySkillID(skillID uint32) (*gameconf.SkillConfDefine, bool) {
	conf, exist := p.cfg.getRawConfig().skillID2Conf[skillID]
	return conf, exist
}
