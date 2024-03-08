package config

import (
	"github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/os"
)

// Config for client.
type Config struct {
	Access        string `yaml:"access" json:"access" toml:"access"`
	client.Config `yaml:",inline" json:",inline" toml:",inline"`
}

// GetAccess for client.
func (c *Config) GetAccess() string {
	return os.GetFromEnv(c.Access)
}
