package config

import (
	"os"
	auth_config "url-shortener/auth/config"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Auth auth_config.Config `yaml:"auth"`
}

func LoadConfig() (*AppConfig, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var cfg AppConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
