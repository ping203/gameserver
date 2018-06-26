package gameconf

import (
	"github.com/sirupsen/logrus"
	"sanguosha.com/games/sgs/framework/util"
)

// GameConfig ...
type GameConfig struct {
	rawConfig *rawConfig
	version   string

	generalConf
	skillConf
	effectConf
	globalConf
}

// Init ...
func (p *GameConfig) Init(path *GameConfigPathNode) error {
	var err error

	p.rawConfig = &rawConfig{}
	err = p.rawConfig.Init(path)
	if err != nil {
		return err
	}

	err = p.load()
	if err != nil {
		return err
	}

	return nil
}

// Reload ...
func (p *GameConfig) Reload() error {
	err := p.rawConfig.Reload()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"time": util.GetCurrentTime().Local().String(),
		}).WithError(err).Error("reload game raw config error")
		return err
	}

	err = p.load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"time": util.GetCurrentTime().Local().String(),
		}).WithError(err).Info("reload game config success")
		return err
	}
	return err
}

func (p *GameConfig) load() error {

	err := p.globalConf.init(p)
	if err != nil {
		return err
	}

	err = p.generalConf.init(p)
	if err != nil {
		return err
	}

	err = p.skillConf.init(p)
	if err != nil {
		return err
	}

	err = p.effectConf.init(p)
	if err != nil {
		return err
	}

	err = p.afterLoad()
	if err != nil {
		return err
	}

	return nil
}

func (p *GameConfig) afterLoad() error {

	return nil
}

func (p *GameConfig) getRawConfig() *rawConfig {
	return p.rawConfig
}

// SetVersion ...
func (p *GameConfig) SetVersion(version string) {
	p.version = version
}

// GetVersion ...
func (p *GameConfig) GetVersion() string {
	return p.version
}
