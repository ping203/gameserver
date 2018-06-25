package gameconf

import (
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gameconf"
	"sanguosha.com/games/sgs/framework/util"
)

type GameConfigPathNode struct {
	// BaseConfigPath
	BaseConfigPath string `yaml:"base_config_path"`
}

type rawConfigNode struct {
	baseCfg *gameconf.GameBaseConfig
}

type rawConfig struct {
	cfgPath *GameConfigPathNode
	cfgNode *rawConfigNode

	generalID2Conf map[uint32]*gameconf.GeneralConfDefine
	skillID2Conf   map[uint32]*gameconf.SkillConfDefine
	effectID2Conf  map[uint32]*gameconf.SkillEffectConfDefine
}

func (p *rawConfig) getContent(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (p *rawConfig) initRawCfgNode() error {
	p.cfgNode = &rawConfigNode{}
	// base
	baseContent, err := p.getContent(p.cfgPath.BaseConfigPath)
	if err != nil {
		return err
	}
	p.cfgNode.baseCfg = &gameconf.GameBaseConfig{}
	err = proto.UnmarshalText(string(baseContent), p.cfgNode.baseCfg)
	if err != nil {
		return err
	}
	return nil
}

// Init ...
func (p *rawConfig) Init(path *GameConfigPathNode) error {
	p.cfgPath = path

	err := p.initRawCfgNode()
	if err != nil {
		logrus.WithError(err).Error("initRawCfgNode error")
		return err
	}

	err = p.load()
	if err != nil {
		logrus.WithError(err).Error("init game config load error")
		return err
	}

	return nil
}

func (p *rawConfig) load() error {
	// base
	err := p.loadBaseCfg()
	if err != nil {
		logrus.WithError(err).Error("initBaseCfg error")
		return err
	}

	return nil
}

// Reload ...
func (p *rawConfig) Reload() error {
	err := p.initRawCfgNode()
	if err != nil {
		logrus.WithError(err).Error("initRawCfgNode error")
		return err
	}

	err = p.load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"time": util.GetCurrentTime().Local().String(),
		}).WithError(err).Error("reload raw config error")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"time": util.GetCurrentTime().Local().String(),
	}).Info("reload raw config success")
	return nil
}

func (p *rawConfig) loadBaseCfg() error {
	p.generalID2Conf = make(map[uint32]*gameconf.GeneralConfDefine)
	for _, v := range p.cfgNode.baseCfg.GetGeneralConf() {
		p.generalID2Conf[v.GeneralID] = v
	}

	p.skillID2Conf = make(map[uint32]*gameconf.SkillConfDefine)
	for _, v := range p.cfgNode.baseCfg.GetSkillConf() {
		p.skillID2Conf[v.SkillID] = v
	}

	p.effectID2Conf = make(map[uint32]*gameconf.SkillEffectConfDefine)
	for _, v := range p.cfgNode.baseCfg.GetSkillEffectConf() {
		p.effectID2Conf[v.SkillEffectID] = v
	}

	return nil
}
