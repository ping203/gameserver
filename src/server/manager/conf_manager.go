package manager

import (
	"fmt"
	"time"

	"server/config/gameconf"
	"server/logs"

	"sanguosha.com/games/sgs/framework/util"
)

// ConfManager ...
type ConfManager struct {
	paths   *gameconf.GameConfigPathNode
	gameCfg *gameconf.GameConfig

	reloadTime time.Time
}

// Init ...
func (p *ConfManager) Init(paths *gameconf.GameConfigPathNode) {
	p.paths = paths

	p.gameCfg = &gameconf.GameConfig{}
	err := p.gameCfg.Init(paths)
	if err != nil {
		logs.Error("paths %v", paths)
		panic(fmt.Errorf("load game config error, error: %v", err))
	}

	p.gameCfg.SetVersion(p.generateVersion())

	logs.Debug("init game config success")
}

func (p *ConfManager) generateVersion() string {
	return fmt.Sprintf("%d", util.GetCurrentTimestamp())
}

func (p *ConfManager) copy() {

}

// Reload ...
func (p *ConfManager) Reload() error {
	// 缓存旧配置(注意拷贝对象的时候先*在&)
	tmp := *p.gameCfg
	oldGameCfg := &tmp
	// 重载
	err := p.gameCfg.Reload()
	if err != nil {
		logs.Error("reload game config, reload game config error")
		// 还原配置
		p.gameCfg = oldGameCfg
		return err
	}

	p.gameCfg.SetVersion(p.generateVersion())
	logs.Debug("reload game config success %v", util.GetCurrentTime().String())
	return nil
}

// GetConfig ...
func (p *ConfManager) GetConfig() *gameconf.GameConfig {
	return p.gameCfg
}

// GetVersion ...
func (p *ConfManager) GetVersion() string {
	return p.gameCfg.GetVersion()
}
