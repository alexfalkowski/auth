package password

import (
	"context"
	"strconv"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

// Secure password.
type Secure struct {
	generator *password.Generator
}

// NewSecure password.
func NewSecure(generator *password.Generator) *Secure {
	return &Secure{generator: generator}
}

// Generate secure password.
func (s *Secure) Generate(ctx context.Context, length Length) (string, error) {
	if err := length.Valid(); err != nil {
		return "", err
	}

	l := int(length)
	meta.WithAttribute(ctx, "password.generate.length", strconv.Itoa(l))

	return password.Generate(l, 10, 10, false, false)
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
