package service

import (
	"time"

	"github.com/alexfalkowski/auth/key"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func generateJWTToken(params TokenParams) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	k, err := key.PrivateEd25519(params.Ed25519.Private)
	if err != nil {
		return "", err
	}

	t := time.Now()

	claims := &jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: t.Add(params.Service.Duration)},
		ID:        id.String(),
		IssuedAt:  &jwt.NumericDate{Time: t},
		Issuer:    params.Issuer,
		NotBefore: &jwt.NumericDate{Time: t},
		Subject:   params.Service.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	return token.SignedString(k)
}
