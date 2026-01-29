package config

import "time"

type Config struct {
	DBConfig   DBConfig   `yaml:"dsn_cfg"`
	ConnConfig ConnConfig `yaml:"conn_cfg"`
}

type DBConfig struct {
	Database string `yaml:"database"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	Server   string `yaml:"server"`
}

type ConnConfig struct {
	MaxConns          int64         `yaml:"max_conns"`
	MinConns          int64         `yaml:"min_conns"`
	MaxConnLifeTime   time.Duration `yaml:"max_conn_life_time"`
	MaxConnIdleTime   time.Duration `yaml:"max_conn_idle_time"`
	HealthCheckPeriod time.Duration `yaml:"health_check_period"`
}
