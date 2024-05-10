package config

import (
	"github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/os"
)

// Config for client.
type Config struct {
	*client.Config `yaml:",inline" json:",inline" toml:",inline"`
	Access         string `yaml:"access,omitempty" json:"access,omitempty" toml:"access,omitempty"`
}

// GetAccess for client.
func (c *Config) GetAccess() string {
	return os.GetFromEnv(c.Access)
}
