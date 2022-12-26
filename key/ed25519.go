package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
)

// Ed25519 cypher.
type Ed25519 struct {
	privateKey ed25519.PrivateKey
}

// NewEd25519 cypher.
func NewEd25519(privateKey ed25519.PrivateKey) *Ed25519 {
	return &Ed25519{privateKey: privateKey}
}

// Generate key pair with Ed25519.
func (e *Ed25519) Generate() (string, string, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(pub), base64.StdEncoding.EncodeToString(pri), nil
}
