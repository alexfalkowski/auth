package service

import (
	"errors"
	"time"
)

var (
	// ErrInvalidKind for generation of tokens.
	ErrInvalidKind = errors.New("invalid kind")
)

// Generator of tokens.
type Generator struct {
	branca *Branca
	jwt    *JWT
	paseto *Paseto
}

// NewGenerator of tokens.
func NewGenerator(branca *Branca, jwt *JWT, paseto *Paseto) *Generator {
	return &Generator{branca: branca, jwt: jwt, paseto: paseto}
}

// Generate token based on kind.
func (g *Generator) Generate(kind, sub, aud, iss string, exp time.Duration) (string, error) {
	switch kind {
	case "jwt":
		return g.jwt.Generate(sub, aud, iss, exp)
	case "branca":
		return g.branca.Generate(sub, aud, iss, exp)
	case "paseto":
		return g.paseto.Generate(sub, aud, iss, exp)
	}

	return "", ErrInvalidKind
}
