package config

import (
	"os"
	auth_config "url-shortener/auth/config"
	db_config "url-shortener/database/config"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Auth     auth_config.Config `yaml:"auth"`
	Database db_config.Config   `yaml:"database"`
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
