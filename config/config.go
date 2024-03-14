package config

import (
	"github.com/alexfalkowski/auth/casbin"
	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/health"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/server"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Key           key.Config    `yaml:"key,omitempty" json:"key,omitempty" toml:"key,omitempty"`
	Casbin        casbin.Config `yaml:"casbin,omitempty" json:"casbin,omitempty" toml:"casbin,omitempty"`
	Client        client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Server        server.Config `yaml:"server,omitempty" json:"server,omitempty" toml:"server,omitempty"`
	Health        health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
