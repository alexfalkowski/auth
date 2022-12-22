package config

import (
	"github.com/alexfalkowski/auth/health"
	v1 "github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func v1Config(cfg config.Configurator) *v1.Config {
	return &cfg.(*Config).Server.V1
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}
