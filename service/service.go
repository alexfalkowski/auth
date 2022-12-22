package service

import (
	"context"
	"time"

	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// SignToken for service.
func SignToken(ctx context.Context, config *config.Config, svc config.Service) (string, error) {
	ctx = meta.WithAttribute(ctx, "service.sign_token.method", "RS512")

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	k, err := key.Private(ctx, config.Key.Private)
	if err != nil {
		return "", err
	}

	t := time.Now()

	claims := &jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: t.Add(svc.Duration)},
		ID:        id.String(),
		IssuedAt:  &jwt.NumericDate{Time: t},
		Issuer:    config.Issuer,
		NotBefore: &jwt.NumericDate{Time: t},
		Subject:   svc.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	return token.SignedString(k)
}
