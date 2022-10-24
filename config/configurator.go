package config

import (
	"github.com/alexfalkowski/auth/health"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}
