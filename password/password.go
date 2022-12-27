package password

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

// Secure password.
type Secure struct {
}

// NewSecure password.
func NewSecure() *Secure {
	return &Secure{}
}

// Generate secure password.
func (s *Secure) Generate(ctx context.Context) (string, error) {
	meta.WithAttribute(ctx, "password.generate.length", "64")

	return password.Generate(64, 10, 10, false, false)
}

// Hash the password using bcrypt.
func (s *Secure) Hash(ctx context.Context, pass string) (string, error) {
	ctx = meta.WithAttribute(ctx, "password.hash.kind", "bcrypt")
	meta.WithAttribute(ctx, "password.hash.cost", "10")

	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

// Compare using bcrypt.
func (s *Secure) Compare(ctx context.Context, hash, pass string) error {
	meta.WithAttribute(ctx, "password.hash.kind", "bcrypt")

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
