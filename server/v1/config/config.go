package config

import (
	"time"
)

// Config for v1.
type Config struct {
	Issuer   string    `yaml:"issuer"`
	Casbin   Casbin    `yaml:"casbin"`
	Key      Key       `yaml:"key"`
	Secret   Secret    `yaml:"secret"`
	Admins   []Admin   `yaml:"admins"`
	Services []Service `yaml:"services"`
}

// Casbin for v1.
type Casbin struct {
	Model  string `yaml:"model"`
	Policy string `yaml:"policy"`
}

// Key for v1.
type Key struct {
	RSA     KeyPair `yaml:"rsa"`
	Ed25519 KeyPair `yaml:"ed25519"`
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

// Secret for v1.
type Secret struct {
	Branca string `yaml:"branca"`
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
