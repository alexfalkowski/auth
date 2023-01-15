package config

import (
	"github.com/alexfalkowski/auth/health"
	"github.com/alexfalkowski/auth/server"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Server        server.Config `yaml:"server" json:"server" toml:"server"`
	Health        health.Config `yaml:"health" json:"health" toml:"health"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
