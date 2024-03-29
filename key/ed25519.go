package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
)

// Ed25519 for key.
type Ed25519 struct {
	privateKey ed25519.PrivateKey
}

// NewEd25519 for key.
func NewEd25519(privateKey ed25519.PrivateKey) *Ed25519 {
	return &Ed25519{privateKey: privateKey}
}

// Generate key pair with Ed25519.
func (e *Ed25519) Generate() (string, string, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)

	return base64.StdEncoding.EncodeToString(pub), base64.StdEncoding.EncodeToString(pri), err
}

// NewEd25519PublicKey from key.
func NewEd25519PublicKey(cfg *Config) (ed25519.PublicKey, error) {
	return base64.StdEncoding.DecodeString(cfg.Ed25519.Public)
}

// NewEd25519PrivateKey from key.
func NewEd25519PrivateKey(cfg *Config) (ed25519.PrivateKey, error) {
	return base64.StdEncoding.DecodeString(cfg.Ed25519.GetPrivate())
}
