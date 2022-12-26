package key

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"

	"github.com/alexfalkowski/go-service/meta"
)

// RSA cypher.
type RSA struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// NewRSA cypher.
func NewRSA(publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) *RSA {
	return &RSA{publicKey: publicKey, privateKey: privateKey}
}

// Encrypt with RSA OAEP.
func (r *RSA) Encrypt(ctx context.Context, msg string) (string, error) {
	ctx = meta.WithAttribute(ctx, "key.encrypt.kind", "OAEP")
	meta.WithAttribute(ctx, "key.encrypt.hash", "SHA512")

	e, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, r.publicKey, []byte(msg), nil)
	if err != nil {
		return "", err
	}

	return string(e), nil
}

// Decrypt with RSA OAEP.
func (r *RSA) Decrypt(ctx context.Context, cipher string) (string, error) {
	ctx = meta.WithAttribute(ctx, "key.decrypt.kind", "OAEP")
	meta.WithAttribute(ctx, "key.decrypt.hash", "SHA512")

	d, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, r.privateKey, []byte(cipher), nil)
	if err != nil {
		return "", err
	}

	return string(d), nil
}

// Generate key pair with RSA.
func (r *RSA) Generate() (string, string, error) {
	p, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	pub := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&p.PublicKey))
	pri := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(p))

	return pub, pri, nil
}
