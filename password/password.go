package password

import (
	"context"
	"errors"
	"strconv"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/matthewhartstonge/argon2"
	"github.com/sethvargo/go-password/password"
)

const symbols = "~!@#$%^&*()_+-={}|[]<>?,./"

// ErrPasswordMismatch when compared.
var ErrPasswordMismatch = errors.New("password mismatch")

// Secure password.
type Secure struct {
	generator *password.Generator
	argon     argon2.Config
}

// NewSecure password.
func NewSecure(generator *password.Generator) *Secure {
	return &Secure{generator: generator, argon: argon2.DefaultConfig()}
}

// Generate secure password.
func (s *Secure) Generate(ctx context.Context, length Length) (string, error) {
	if err := length.Valid(); err != nil {
		return "", err
	}

	l := int(length)
	meta.WithAttribute(ctx, "passwordGenerateLength", meta.String(strconv.Itoa(l)))

	g, err := password.NewGenerator(&password.GeneratorInput{Symbols: symbols})
	if err != nil {
		return "", err
	}

	return g.Generate(l, 10, 10, false, false)
}

// Hash the password.
func (s *Secure) Hash(ctx context.Context, pass string) (string, error) {
	meta.WithAttribute(ctx, "passwordHashKind", meta.String("Argon2id"))

	h, err := s.argon.HashEncoded([]byte(pass))

	return string(h), err
}

// Compare the password with the hash.
func (s *Secure) Compare(ctx context.Context, hash, pass string) error {
	meta.WithAttribute(ctx, "passwordHashKind", meta.String("Argon2id"))

	ok, err := argon2.VerifyEncoded([]byte(pass), []byte(hash))
	if err != nil {
		return err
	}

	if !ok {
		return ErrPasswordMismatch
	}

	return nil
}
