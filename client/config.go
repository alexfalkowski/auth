package client

import (
	v1 "github.com/alexfalkowski/auth/client/v1/config"
)

// IsEnabled for client.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for client.
type Config struct {
	V1 *v1.Config `yaml:"v1,omitempty" json:"v1,omitempty" toml:"v1,omitempty"`
}
