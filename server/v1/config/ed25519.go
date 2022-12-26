package config

import (
	"crypto/ed25519"
	"encoding/base64"
)

// PrivateRSA from key.
func NewEd25519PrivateKey(cfg *Config) (ed25519.PrivateKey, error) {
	return base64.StdEncoding.DecodeString(cfg.Key.Ed25519.Private)
}
