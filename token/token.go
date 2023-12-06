package token

import (
	"errors"
	"time"
)

var (
	// ErrInvalidKind for generation of tokens.
	ErrInvalidKind = errors.New("invalid kind")

	// ErrInvalidAlgorithm for service.
	ErrInvalidAlgorithm = errors.New("invalid algorithm")

	// ErrInvalidIssuer for service.
	ErrInvalidIssuer = errors.New("invalid issuer")

	// ErrInvalidAudience for service.
	ErrInvalidAudience = errors.New("invalid audience")

	// ErrInvalidTime for service.
	ErrInvalidTime = errors.New("invalid time")
)

// Token generator.
type Token struct {
	jwt    *JWT
	paseto *Paseto
}

// NewToken generator.
func NewToken(jwt *JWT, paseto *Paseto) *Token {
	return &Token{jwt: jwt, paseto: paseto}
}

// Generate token based on kind.
func (t *Token) Generate(kind, sub, aud, iss string, exp time.Duration) (string, error) {
	switch kind {
	case "jwt":
		return t.jwt.Generate(sub, aud, iss, exp)
	case "paseto":
		return t.paseto.Generate(sub, aud, iss, exp)
	}

	return "", ErrInvalidKind
}

// Verify token based on kind.
func (t *Token) Verify(token, kind, aud, iss string) (string, error) {
	switch kind {
	case "jwt":
		return t.jwt.Verify(token, aud, iss)
	case "paseto":
		return t.paseto.Verify(token, aud, iss)
	}

	return "", ErrInvalidKind
}
