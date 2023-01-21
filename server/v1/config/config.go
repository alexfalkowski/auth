package config

import (
	"time"
)

// Config for v1.
type Config struct {
	Issuer   string    `yaml:"issuer" json:"issuer" toml:"issuer"`
	Casbin   Casbin    `yaml:"casbin" json:"casbin" toml:"casbin"`
	Admins   []Admin   `yaml:"admins" json:"admins" toml:"admins"`
	Services []Service `yaml:"services" json:"services" toml:"services"`
}

// Casbin for v1.
type Casbin struct {
	Model  string `yaml:"model" json:"model" toml:"model"`
	Policy string `yaml:"policy" json:"policy" toml:"policy"`
}

// Admin for v1.
type Admin struct {
	ID   string `yaml:"id"`
	Hash string `yaml:"hash"`
}

// Service for v1.
type Service struct {
	ID       string        `yaml:"id"`
	Hash     string        `yaml:"hash"`
	Duration time.Duration `yaml:"duration"`
}
