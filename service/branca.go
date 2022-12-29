package service

import (
	"encoding/json"
	"time"

	"github.com/essentialkaos/branca"
)

// BrancaToken for service.
type BrancaToken struct {
	Subject  string `json:"sub"`
	Audience string `json:"aud"`
	Issuer   string `json:"iss"`
}

// Branca service to generate tokens.
type Branca struct {
	branca *branca.Branca
}

// NewBranca service to generate tokens.
func NewBranca(branca *branca.Branca) *Branca {
	return &Branca{branca: branca}
}

// Generate Branca token.
func (b *Branca) Generate(sub, aud, iss string, exp time.Duration) (string, error) {
	t := time.Now()
	b.branca.SetTTL(uint32(t.Add(exp).Unix()))

	bt := BrancaToken{Subject: sub, Audience: aud, Issuer: iss}

	by, err := json.Marshal(bt)
	if err != nil {
		return "", err
	}

	return b.branca.EncodeToString(by)
}

// Verify Branca token.
func (b *Branca) Verify(token, iss string) (string, string, error) {
	to, err := b.branca.DecodeString(token)
	if err != nil {
		return "", "", err
	}

	if b.branca.IsExpired(to) {
		return "", "", ErrInvalidTime
	}

	var bt BrancaToken
	if err := json.Unmarshal(to.Payload(), &bt); err != nil {
		return "", "", err
	}

	if bt.Issuer != iss {
		return "", "", ErrInvalidIssuer
	}

	return bt.Subject, bt.Audience, nil
}
