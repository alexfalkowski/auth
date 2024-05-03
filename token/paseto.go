package token

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/alexfalkowski/auth/key"
	"github.com/google/uuid"
)

// Paseto token.
type Paseto struct {
	ed *key.Ed25519
}

// NewPaseto token.
func NewPaseto(ed *key.Ed25519) *Paseto {
	return &Paseto{ed: ed}
}

// Generate Paseto token.
func (p *Paseto) Generate(sub, aud, iss string, exp time.Duration) (string, error) {
	t := time.Now()
	token := paseto.NewToken()

	token.SetJti(uuid.NewString())
	token.SetIssuedAt(t)
	token.SetNotBefore(t)
	token.SetExpiration(t.Add(exp))
	token.SetIssuer(iss)
	token.SetSubject(sub)
	token.SetAudience(aud)

	s, err := paseto.NewV4AsymmetricSecretKeyFromBytes(p.ed.PrivateKey())
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

	s, err := paseto.NewV4AsymmetricPublicKeyFromBytes(p.ed.PublicKey())
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
