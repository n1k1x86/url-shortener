package config

import "time"

type Config struct {
	Base         Base         `yaml:"base"`
	RefreshToken RefreshToken `yaml:"refresh"`
	AccessToken  AccessToken  `yaml:"access"`
}

type Base struct {
	Issuer   string `yaml:"issuer"`
	Audience string `yaml:"audience"`
}

type AccessToken struct {
	ExpiredAfter time.Duration `yaml:"expired_after"`
	Secret       string        `yaml:"secret"`
}

type RefreshToken struct {
	ExpiredAfter time.Duration `yaml:"expired_after"`
	Secret       string        `yaml:"secret"`
}
