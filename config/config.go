package config

import (
	"github.com/alexfalkowski/auth/casbin"
	"github.com/alexfalkowski/auth/client"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/auth/health"
	"github.com/alexfalkowski/auth/server"
	v1s "github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfig for config.
func NewConfig(i *cmd.InputConfig) (*Config, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

// Config for the service.
type Config struct {
	Casbin         *casbin.Config `yaml:"casbin,omitempty" json:"casbin,omitempty" toml:"casbin,omitempty"`
	Client         *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Server         *server.Config `yaml:"server,omitempty" json:"server,omitempty" toml:"server,omitempty"`
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func casbinConfig(cfg *Config) *casbin.Config {
	return cfg.Casbin
}

func v1ServerConfig(cfg *Config) *v1s.Config {
	if !server.IsEnabled(cfg.Server) {
		return nil
	}

	return cfg.Server.V1
}

func v1ClientConfig(cfg *Config) *v1c.Config {
	if !client.IsEnabled(cfg.Client) {
		return nil
	}

	return cfg.Client.V1
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}
