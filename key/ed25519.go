package key

import (
	"github.com/alexfalkowski/go-service/crypto/ed25519"
)

// Ed25519 for key.
type Ed25519 struct {
	algo ed25519.Algo
}

// NewEd25519 for key.
func NewEd25519(algo ed25519.Algo) *Ed25519 {
	return &Ed25519{algo: algo}
}

// Generate key pair with Ed25519.
func (e *Ed25519) Generate() (string, string, error) {
	pub, pri, err := ed25519.Generate()

	return pub, pri, err
}

// Algo for Ed25519.
func (e *Ed25519) Algo() ed25519.Algo {
	return e.algo
}
