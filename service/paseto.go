package service

import (
	"crypto/ed25519"
	"time"

	"aidanwoods.dev/go-paseto"
)

// Paseto service.
type Paseto struct {
	privateKey ed25519.PrivateKey
}

// NewPaseto service.
func NewPaseto(privateKey ed25519.PrivateKey) *Paseto {
	return &Paseto{privateKey: privateKey}
}

// Generate Paseto token.
func (p *Paseto) Generate(sub, iss string, exp time.Duration) (string, error) {
	t := time.Now()
	token := paseto.NewToken()

	token.SetIssuedAt(t)
	token.SetNotBefore(t)
	token.SetExpiration(t.Add(exp))
	token.SetIssuer(iss)
	token.SetSubject(sub)

	s, err := paseto.NewV4AsymmetricSecretKeyFromBytes(p.privateKey)
	if err != nil {
		return "", err
	}

	return token.V4Sign(s, nil), nil
}
