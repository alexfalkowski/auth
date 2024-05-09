package key

import (
	"github.com/alexfalkowski/go-service/crypto/rsa"
)

// RSA cipher.
type RSA struct {
	algo rsa.Algo
}

// NewRSA cipher.
func NewRSA(algo rsa.Algo) *RSA {
	return &RSA{algo: algo}
}

// Generate key pair with RSA.
func (r *RSA) Generate() (string, string, error) {
	return rsa.Generate()
}

// Encrypt with RSA OAEP.
func (r *RSA) Encrypt(msg string) (string, error) {
	return r.algo.Encrypt(msg)
}

// Decrypt with RSA OAEP.
func (r *RSA) Decrypt(msg string) (string, error) {
	return r.algo.Decrypt(msg)
}
