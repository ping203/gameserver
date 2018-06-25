package gameconf

import "github.com/wenxiu2199/gameserver/src/server/gameproto/gameconf"

type generalConf struct {
	cfg *GameConfig
}

func (p *generalConf) init(cfg *GameConfig) error {
	p.cfg = cfg

	return nil
}

func (p *generalConf) GetGeneralConfByGeneralID(generalID uint32) (*gameconf.GeneralConfDefine, bool) {
	conf, exist := p.cfg.getRawConfig().generalID2Conf[generalID]
	return conf, exist
}
