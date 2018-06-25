package gameconf

import "github.com/wenxiu2199/gameserver/src/server/gameproto/gameconf"

type effectConf struct {
	cfg *GameConfig
}

func (p *effectConf) init(cfg *GameConfig) error {
	p.cfg = cfg

	return nil
}

func (p *effectConf) GetEffectConfByEffectID(effectID uint32) (*gameconf.SkillEffectConfDefine, bool) {
	conf, exist := p.cfg.getRawConfig().effectID2Conf[effectID]
	return conf, exist
}
