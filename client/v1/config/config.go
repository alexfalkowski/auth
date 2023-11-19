package config

import (
	"time"

	"github.com/alexfalkowski/go-service/os"
)

// Config for client.
type Config struct {
	Host    string        `yaml:"host" json:"host" toml:"host"`
	Timeout time.Duration `yaml:"timeout" json:"timeout" toml:"timeout"`
	Access  string        `yaml:"access" json:"access" toml:"access"`
}

// GetAccess for client.
func (c *Config) GetAccess() string {
	return os.GetFromEnv(c.Access)
}
