package config

import (
	"github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/os"
)

// IsEnabled config.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

type (

	// Access for config.
	Access string

	// Config for client.
	Config struct {
		*client.Config `yaml:",inline" json:",inline" toml:",inline"`
		Access         Access `yaml:"access,omitempty" json:"access,omitempty" toml:"access,omitempty"`
	}
)

// GetAccess for client.
func (c *Config) GetAccess() (string, error) {
	return os.ReadFile(string(c.Access))
}
