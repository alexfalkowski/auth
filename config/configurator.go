package config

import (
	"github.com/alexfalkowski/auth/casbin"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/auth/health"
	"github.com/alexfalkowski/auth/key"
	v1s "github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfigurator for config.
func NewConfigurator(i *cmd.InputConfig) (config.Configurator, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

func cfg(cfg config.Configurator) *Config {
	return cfg.(*Config)
}

func casbinConfig(cfg config.Configurator) *casbin.Config {
	return &cfg.(*Config).Casbin
}

func keyConfig(cfg config.Configurator) *key.Config {
	return &cfg.(*Config).Key
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
