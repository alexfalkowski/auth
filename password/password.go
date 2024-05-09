package password

import (
	"github.com/alexfalkowski/go-service/crypto/argon2"
	"github.com/alexfalkowski/go-service/crypto/rand"
)

// Secure password.
type Secure struct {
	algo argon2.Algo
}

// NewSecure password.
func NewSecure(algo argon2.Algo) *Secure {
	return &Secure{algo: algo}
}

// Generate secure password.
func (s *Secure) Generate(length Length) (string, error) {
	if err := length.Valid(); err != nil {
		return "", err
	}

	return rand.GenerateString(uint32(length))
}

// Hash the password.
func (s *Secure) Hash(pass string) (string, error) {
	return s.algo.Generate(pass)
}

// Compare the password with the hash.
func (s *Secure) Compare(hash, pass string) error {
	return s.algo.Compare(hash, pass)
}
