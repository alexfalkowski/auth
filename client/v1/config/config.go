package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/alexfalkowski/go-service/client"
)

type (

	// Access for config.
	Access string

	// Config for client.
	Config struct {
		*client.Config `yaml:",inline" json:",inline" toml:",inline"`
		Access         Access `yaml:"access,omitempty" json:"access,omitempty" toml:"access,omitempty"`
	}
)

// GetKey for client.
func (c *Config) GetAccess() (string, error) {
	k, err := os.ReadFile(filepath.Clean(string(c.Access)))

	return strings.TrimSpace(string(k)), err
}
