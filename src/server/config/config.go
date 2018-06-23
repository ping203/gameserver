package config

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
	"sanguosha.com/games/sgs/config/gameconf"
)

// RedisNode 一个RedisNode节点.
type RedisNode struct {
	// Addr 服务器地址
	Addr     string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Size     int    `yaml:"size"`
}

// FilterNode 一个FilterNode节点.
type FilterNode struct {
	// Address 服务器地址
	Address string `yaml:"address"`
}

// GeoIPNode 一个FilterNode节点.
type GeoIPNode struct {
	// FilePath 数据文件
	FilePath string `yaml:"file_path"`
	// Lang 显示语言
	Lang string `yaml:"lang"`
}

// NatsNode 一个NatsNode节点.
type NatsNode struct {
	// Cluster 服务器集群名称
	Cluster string `yaml:"cluster"`
	// Address 服务器地址
	Address string `yaml:"address"`
	// OffsetPath 偏移量文件.
	OffsetPath string `yaml:"offset_path"`
}

type DBSPool struct {
	MaxOpen     int `yaml:"maxOpen"`     //max open conns
	MaxIdle     int `yaml:"maxIdle"`     //max idle conns
	MaxLifetime int `yaml:"maxLifetime"` //conn max lifetime
}

// DBSNode 一个DBS节点.
type DBSNode struct {
	Address string  `yaml:"address"`
	Pool    DBSPool `yaml:"pool"`
}

// RouterNode ...
type RouterNode struct {
	ServerID uint32 `yaml:"server_id"`
	Address  string `yaml:"address"`
}

// Config ...
type Config struct {
	Debug bool `yaml:"debug"`

	NetConfigFile string             `yaml:"netconfig"`
	Redis         *RedisNode         `yaml:"redis"`
	Filter        *FilterNode        `yaml:"filter"`
	GeoIP         *GeoIPNode         `yaml:"geoIPDatabase"`
	Nats          *NatsNode          `yaml:"nats"`
	DBS           map[string]DBSNode `yaml:"dbs"`

	NetEaseAppKey    string `yaml:"netease_appkey"`
	NetEaseAppSecret string `yaml:"netease_appsecret"`

	GameCfgPath *gameconf.GameConfigPathNode `yaml:"game_config_path"`

	CasbinCfgPath *gameconf.CasbinConfigPathNode `yaml:"casbin_config_path"`
}

var appCfg *Config

func parseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	if cfg.NetConfigFile == "" {
		cfg.NetConfigFile = "netconfig.json"
	}
	if cfg.Redis == nil {
		return nil, errors.New("no redis config")
	}
	if cfg.GeoIP == nil {
		return nil, errors.New("no geoIP config")
	}
	if len(cfg.DBS) == 0 {
		return nil, errors.New("no dbs config")
	}

	if cfg.GameCfgPath == nil {
		return nil, errors.New("no game file path config")
	}

	return &cfg, nil
}

// ParseConfigFile ...
func ParseConfigFile(fileName string) (*Config, error) {
	abs, err := filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(abs)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, err
	}

	cfg, err := parseConfigData(data)
	if err != nil {
		return nil, err
	}

	cfg.GeoIP.FilePath = filepath.Join(dir, cfg.GeoIP.FilePath)
	cfg.Nats.OffsetPath = filepath.Join(dir, cfg.Nats.OffsetPath)
	cfg.GameCfgPath.BaseConfigPath = filepath.Join(dir, cfg.GameCfgPath.BaseConfigPath)
	cfg.GameCfgPath.ItemConfigPath = filepath.Join(dir, cfg.GameCfgPath.ItemConfigPath)
	cfg.GameCfgPath.ModeConfigPath = filepath.Join(dir, cfg.GameCfgPath.ModeConfigPath)
	cfg.GameCfgPath.DropConfigPath = filepath.Join(dir, cfg.GameCfgPath.DropConfigPath)
	cfg.GameCfgPath.GuildConfigPath = filepath.Join(dir, cfg.GameCfgPath.GuildConfigPath)

	appCfg = cfg

	return cfg, nil
}

// GetConfig ...
func GetConfig() (*Config, error) {
	if appCfg == nil {
		return nil, errors.New("init server config first")
	}
	return appCfg, nil
}
