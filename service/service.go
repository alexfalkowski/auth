package service

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

// Service of tokens.
type Service struct {
	jwt    *JWT
	paseto *Paseto
}

// NewService of tokens.
func NewService(jwt *JWT, paseto *Paseto) *Service {
	return &Service{jwt: jwt, paseto: paseto}
}

// Generate token based on kind.
func (s *Service) Generate(kind, sub, aud, iss string, exp time.Duration) (string, error) {
	switch kind {
	case "jwt":
		return s.jwt.Generate(sub, aud, iss, exp)
	case "paseto":
		return s.paseto.Generate(sub, aud, iss, exp)
	}

	return "", ErrInvalidKind
}

// Verify token based on kind.
func (s *Service) Verify(token, kind, aud, iss string) (string, error) {
	switch kind {
	case "jwt":
		return s.jwt.Verify(token, aud, iss)
	case "paseto":
		return s.paseto.Verify(token, aud, iss)
	}

	return "", ErrInvalidKind
}
