package service

import (
	"time"

	"github.com/essentialkaos/branca"
)

// BrancaSecret service.
type BrancaSecret string

// Branca service.
type Branca struct {
	secret BrancaSecret
}

// NewBranca service.
func NewBranca(secret BrancaSecret) *Branca {
	return &Branca{secret: secret}
}

// Generate Branca token.
func (b *Branca) Generate(sub, iss string, exp time.Duration) (string, error) {
	brc, err := branca.NewBranca([]byte(b.secret))
	if err != nil {
		return "", err
	}

	t := time.Now()
	brc.SetTTL(uint32(t.Add(exp).Unix()))

	return brc.EncodeToString([]byte(sub))
}
