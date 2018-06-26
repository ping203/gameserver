package gameconf

import (
	"errors"
	"fmt"

	"github.com/wenxiu2199/gameserver/src/server/gameproto/gameconf"
)

type globalConf struct {
	cfg *GameConfig

	conf *gameconf.GlobalconfDefine
}

func (p *globalConf) init(cfg *GameConfig) error {
	p.cfg = cfg

	conf := cfg.getRawConfig().cfgNode.baseCfg.GetGlobalconf()
	if len(conf) == 0 {
		return errors.New("init global config error")
	}

	p.conf = conf[0]

	err := p.checkGlobal()
	if err != nil {
		return fmt.Errorf("init global config error: %v", err)
	}

	return nil
}

// GetGlobalConfig ...
func (p *globalConf) GetGlobalConfig() *gameconf.GlobalconfDefine {
	return p.conf
}

func (p *globalConf) logError(method, key, format string, args ...interface{}) error {
	format = fmt.Sprintf("method: %s, key: %s, %s", method, key, format)
	return fmt.Errorf(format, args...)
}

// 检查配置
func (p *globalConf) checkGlobal() error {
	return nil
}
