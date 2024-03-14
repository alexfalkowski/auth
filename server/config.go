package server

import (
	v1 "github.com/alexfalkowski/auth/server/v1/config"
)

// Config for server.
type Config struct {
	V1 v1.Config `yaml:"v1,omitempty" json:"v1,omitempty" toml:"v1,omitempty"`
}
