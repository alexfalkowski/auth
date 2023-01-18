package config

import (
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/auth/health"
	v1s "github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func v1ServerConfig(cfg config.Configurator) *v1s.Config {
	return &cfg.(*Config).Server.V1
}

func v1ClientConfig(cfg config.Configurator) *v1c.Config {
	return &cfg.(*Config).Client.V1
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}
