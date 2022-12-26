package key

import (
	"errors"
)

var (
	// ErrInvalidKind for key.
	ErrInvalidKind = errors.New("invalid kind")
)

// Generator of key pairs.
type Generator struct {
	rsa    *RSA
	ed2551 *Ed25519
}

// NewGenerator of key pairs.
func NewGenerator(rsa *RSA, ed2551 *Ed25519) *Generator {
	return &Generator{rsa: rsa, ed2551: ed2551}
}

// Generate key pair based on kind.
func (g *Generator) Generate(kind string) (string, string, error) {
	switch kind {
	case "rsa":
		return g.rsa.Generate()
	case "ed25519":
		return g.ed2551.Generate()
	}

	return "", "", ErrInvalidKind
}
