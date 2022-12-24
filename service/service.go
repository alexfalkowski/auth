package service

import (
	"errors"
	"time"
)

// Service for generation of tokens.
type Service struct {
	ID       string
	Hash     string
	Duration time.Duration
}

// KeyPair for generation of tokens.
type KeyPair struct {
	Public  string
	Private string
}

// TokenParams for generation.
type TokenParams struct {
	RSA     KeyPair
	Ed25519 KeyPair
	Branca  string
	Service Service
	Issuer  string
	Kind    string
}

var (
	// ErrInvalidKind for generation of tokens.
	ErrInvalidKind = errors.New("invalid kind")
)

// GenerateToken for service.
func GenerateToken(params TokenParams) (string, error) {
	switch params.Kind {
	case "jwt":
		return generateJWTToken(params)
	case "branca":
		return generateBrancaToken(params)
	case "paseto":
		return generatePasetoToken(params)
	}

	return "", ErrInvalidKind
}
