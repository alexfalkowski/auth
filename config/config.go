package config

import (
	"github.com/alexfalkowski/auth/health"
	"github.com/alexfalkowski/auth/server"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Server        server.Config `yaml:"server"`
	Health        health.Config `yaml:"health"`
	config.Config `yaml:",inline"`
}
