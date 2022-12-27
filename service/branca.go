package service

import (
	"encoding/json"
	"time"

	"github.com/essentialkaos/branca"
)

// BrancaSecret for service.
type BrancaSecret string

// BrancaToken for service.
type BrancaToken struct {
	Subject  string `json:"sub"`
	Audience string `json:"aud"`
	Issuer   string `json:"iss"`
}

// Branca service to generate tokens.
type Branca struct {
	secret BrancaSecret
}

// NewBranca service to generate tokens.
func NewBranca(secret BrancaSecret) *Branca {
	return &Branca{secret: secret}
}

// Generate Branca token.
func (b *Branca) Generate(sub, aud, iss string, exp time.Duration) (string, error) {
	brc, err := branca.NewBranca([]byte(b.secret))
	if err != nil {
		return "", err
	}

	t := time.Now()
	brc.SetTTL(uint32(t.Add(exp).Unix()))

	to := BrancaToken{Subject: sub, Audience: aud, Issuer: iss}

	by, err := json.Marshal(to)
	if err != nil {
		return "", err
	}

	return brc.EncodeToString(by)
}
