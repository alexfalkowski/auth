package config

import (
	"time"
)

// Config for v1.
type Config struct {
	Issuer   string    `yaml:"issuer"`
	Key      Key       `yaml:"key"`
	Admins   []Admin   `yaml:"admins"`
	Services []Service `yaml:"services"`
}

// Key for v1.
type Key struct {
	Public  string `yaml:"public"`
	Private string `yaml:"private"`
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
