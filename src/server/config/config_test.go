package config

import "testing"

func TestConfig(t *testing.T) {
	cfg, err := ParseConfigFile("./config.example.yaml")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", *cfg)
	}
}
