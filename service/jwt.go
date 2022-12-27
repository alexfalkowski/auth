package service

import (
	"crypto/ed25519"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// JWT service.
type JWT struct {
	privateKey ed25519.PrivateKey
}

// NewJWT service.
func NewJWT(privateKey ed25519.PrivateKey) *JWT {
	return &JWT{privateKey: privateKey}
}

// Generate JWT token.
func (j *JWT) Generate(sub, aud, iss string, exp time.Duration) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	t := time.Now()

	claims := &jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: t.Add(exp)},
		ID:        id.String(),
		IssuedAt:  &jwt.NumericDate{Time: t},
		Issuer:    iss,
		NotBefore: &jwt.NumericDate{Time: t},
		Subject:   sub,
		Audience:  []string{aud},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	return token.SignedString(j.privateKey)
}
