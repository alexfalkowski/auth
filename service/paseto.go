package service

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/alexfalkowski/auth/key"
)

func generatePasetoToken(params TokenParams) (string, error) {
	t := time.Now()
	token := paseto.NewToken()

	token.SetIssuedAt(t)
	token.SetNotBefore(t)
	token.SetExpiration(t.Add(params.Service.Duration))
	token.SetIssuer(params.Issuer)

	b, err := key.PrivateEd25519(params.Ed25519.Private)
	if err != nil {
		return "", err
	}

	s, err := paseto.NewV4AsymmetricSecretKeyFromBytes(b)
	if err != nil {
		return "", err
	}

	return token.V4Sign(s, nil), nil
}
