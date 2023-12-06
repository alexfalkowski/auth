package token

import (
	"crypto/ed25519"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

// Paseto token.
type Paseto struct {
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

// NewPaseto token.
func NewPaseto(publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey) *Paseto {
	return &Paseto{publicKey: publicKey, privateKey: privateKey}
}

// Generate Paseto token.
func (p *Paseto) Generate(sub, aud, iss string, exp time.Duration) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	t := time.Now()
	token := paseto.NewToken()

	token.SetJti(id.String())
	token.SetIssuedAt(t)
	token.SetNotBefore(t)
	token.SetExpiration(t.Add(exp))
	token.SetIssuer(iss)
	token.SetSubject(sub)
	token.SetAudience(aud)

	s, err := paseto.NewV4AsymmetricSecretKeyFromBytes(p.privateKey)
	if err != nil {
		return "", err
	}

	return token.V4Sign(s, nil), nil
}

// Verify Paseto token.
func (p *Paseto) Verify(token, aud, iss string) (string, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.IssuedBy(iss))
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))
	parser.AddRule(paseto.ForAudience(aud))

	s, err := paseto.NewV4AsymmetricPublicKeyFromBytes(p.publicKey)
	if err != nil {
		return "", err
	}

	to, err := parser.ParseV4Public(s, token, nil)
	if err != nil {
		return "", err
	}

	sub, err := to.GetSubject()
	if err != nil {
		return "", err
	}

	return sub, nil
}
