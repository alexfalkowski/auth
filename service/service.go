package service

import (
	"context"
	"time"

	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/essentialkaos/branca"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// GenerateToken for service.
func GenerateToken(ctx context.Context, config *config.Config, svc config.Service, kind string) (string, error) {
	switch kind {
	case "jwt":
		return generateJWTToken(ctx, config, svc)
	case "branca":
		return generateBrancaToken(ctx, config, svc)
	default:
		return generateJWTToken(ctx, config, svc)
	}
}

func generateJWTToken(ctx context.Context, config *config.Config, svc config.Service) (string, error) {
	ctx = meta.WithAttribute(ctx, "service.generate.kind", "jwt")
	ctx = meta.WithAttribute(ctx, "service.generate.method", "RS512")

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	k, err := key.Private(ctx, config.Key.RSA.Private)
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

func generateBrancaToken(ctx context.Context, config *config.Config, svc config.Service) (string, error) {
	meta.WithAttribute(ctx, "service.generate.kind", "branca")

	brc, err := branca.NewBranca([]byte(config.Secret.Branca))
	if err != nil {
		return "", err
	}

	t := time.Now()
	brc.SetTTL(uint32(t.Add(svc.Duration).Unix()))

	return brc.EncodeToString([]byte(svc.ID))
}
