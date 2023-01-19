package config

import (
	"time"
)

// Config for v1.
type Config struct {
	Issuer   string    `yaml:"issuer" json:"issuer" toml:"issuer"`
	Casbin   Casbin    `yaml:"casbin" json:"casbin" toml:"casbin"`
	Key      Key       `yaml:"key" json:"key" toml:"key"`
	Admins   []Admin   `yaml:"admins" json:"admins" toml:"admins"`
	Services []Service `yaml:"services" json:"services" toml:"services"`
}

// Casbin for v1.
type Casbin struct {
	Model  string `yaml:"model" json:"model" toml:"model"`
	Policy string `yaml:"policy" json:"policy" toml:"policy"`
}

// Key for v1.
type Key struct {
	RSA     KeyPair `yaml:"rsa" json:"rsa" toml:"rsa"`
	Ed25519 KeyPair `yaml:"ed25519" json:"ed25519" toml:"ed25519"`
}

// Pair from kind.
func (k *Key) Pair(kind string) *KeyPair {
	switch kind {
	case "rsa":
		return &k.RSA
	case "ed25519":
		return &k.Ed25519
	default:
		return nil
	}
}

// RSA for v1.
type KeyPair struct {
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
