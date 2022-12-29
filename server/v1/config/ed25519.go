package config

import (
	"crypto/ed25519"
	"encoding/base64"
)

// NewEd25519PublicKey from key.
func NewEd25519PublicKey(cfg *Config) (ed25519.PublicKey, error) {
	return base64.StdEncoding.DecodeString(cfg.Key.Ed25519.Public)
}

// NewEd25519PrivateKey from key.
func NewEd25519PrivateKey(cfg *Config) (ed25519.PrivateKey, error) {
	return base64.StdEncoding.DecodeString(cfg.Key.Ed25519.Private)
}
