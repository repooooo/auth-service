package config

import (
	"time"
)

type Config struct {
	Env  string `yaml:"env" env-required:"true"`
	GRPC GRPCConfig
	DSN  string `yaml:"dsn" env-required:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

func New() *Config {
	return &Config{}
}
