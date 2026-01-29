package config

import "time"

type Config struct {
	DBConfig   DBConfig   `yaml:"dsn_cfg"`
	ConnConfig ConnConfig `yaml:"conn_cfg"`
	PoolConfig PoolConfig `yaml:"pool_cfg"`
}

type DBConfig struct {
	Database string `yaml:"database"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	Server   string `yaml:"server"`
}

type ConnConfig struct {
	MaxConns          int32         `yaml:"max_conns"`
	MinConns          int32         `yaml:"min_conns"`
	MaxConnLifeTime   time.Duration `yaml:"max_conn_life_time"`
	MaxConnIdleTime   time.Duration `yaml:"max_conn_idle_time"`
	HealthCheckPeriod time.Duration `yaml:"health_check_period"`
}

type PoolConfig struct {
	PingTimeout    time.Duration `yaml:"ping_timeout"`
	AcquireTimeout time.Duration `yaml:"acquire_timeout"`
}
