package key

import (
	"crypto/ed25519"

	ce "github.com/alexfalkowski/go-service/crypto/ed25519"
)

// Ed25519 for key.
type Ed25519 struct {
	config *ce.Config
}

// NewEd25519 for key.
func NewEd25519(config *ce.Config) *Ed25519 {
	return &Ed25519{config: config}
}

// Generate key pair with Ed25519.
func (e *Ed25519) Generate() (string, string, error) {
	return ce.Generate()
}

// PublicKey for Ed25519.
func (e *Ed25519) PublicKey() (ed25519.PublicKey, error) {
	return e.config.PublicKey()
}

// PrivateKey for Ed25519.
func (e *Ed25519) PrivateKey() (ed25519.PrivateKey, error) {
	return e.config.PrivateKey()
}
