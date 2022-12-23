package key

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"

	"github.com/alexfalkowski/go-service/meta"
)

// Public from key.
func Public(ctx context.Context, key string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(key))

	return x509.ParsePKCS1PublicKey(block.Bytes)
}

// Private from key.
func Private(ctx context.Context, key string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(key))

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// Generate public and private key or error.
func Generate(ctx context.Context) (string, string, error) {
	meta.WithAttribute(ctx, "key.generate.bits", "4096")

	pub, pri, err := generateKeyPair(4096)
	if err != nil {
		return "", "", err
	}

	return exportPublicKeyToPEM(pub), exportPrivateKeyToPEM(pri), nil
}

// Encrypt with public key.
func Encrypt(ctx context.Context, key, pass string) (string, error) {
	ctx = meta.WithAttribute(ctx, "key.encrypt.kind", "OAEP")
	meta.WithAttribute(ctx, "key.encrypt.hash", "sha512")

	k, err := Public(ctx, key)
	if err != nil {
		return "", err
	}

	e, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, k, []byte(pass), nil)
	if err != nil {
		return "", err
	}

	return string(e), nil
}

// Encrypt with private key.
func Decrypt(ctx context.Context, key, cipher string) (string, error) {
	ctx = meta.WithAttribute(ctx, "key.decrypt.kind", "OAEP")
	meta.WithAttribute(ctx, "key.decrypt.hash", "sha512")

	k, err := Private(ctx, key)
	if err != nil {
		return "", err
	}

	d, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, k, []byte(cipher), nil)
	if err != nil {
		return "", err
	}

	return string(d), nil
}

func generateKeyPair(bits int) (*rsa.PublicKey, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	return &privateKey.PublicKey, privateKey, nil
}

func exportPublicKeyToPEM(key *rsa.PublicKey) string {
	return string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(key),
		},
	))
}

func exportPrivateKeyToPEM(key *rsa.PrivateKey) string {
	return string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	))
}
