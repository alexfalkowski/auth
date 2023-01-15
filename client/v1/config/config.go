package config

import (
	"time"
)

// Config for client.
type Config struct {
	Host    string        `yaml:"host" json:"host" toml:"host"`
	Timeout time.Duration `yaml:"timeout" json:"timeout" toml:"timeout"`
	Admin   string        `yaml:"admin" json:"admin" toml:"admin"`
	Access  string        `yaml:"access" json:"access" toml:"access"`
}
