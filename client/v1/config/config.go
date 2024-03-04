package config

import (
	"time"

	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/retry"
	"github.com/alexfalkowski/go-service/security"
)

// Config for client.
type Config struct {
	Host      string          `yaml:"host" json:"host" toml:"host"`
	Security  security.Config `yaml:"security" json:"security" toml:"security"`
	Timeout   time.Duration   `yaml:"timeout" json:"timeout" toml:"timeout"`
	Retry     retry.Config    `yaml:"retry" json:"retry" toml:"retry"`
	UserAgent string          `yaml:"user_agent" json:"user_agent" toml:"user_agent"`
	Access    string          `yaml:"access" json:"access" toml:"access"`
}

// GetAccess for client.
func (c *Config) GetAccess() string {
	return os.GetFromEnv(c.Access)
}
