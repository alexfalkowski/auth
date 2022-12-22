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

// Generate public and private key or error.
func Generate(ctx context.Context) (string, string, error) {
	meta.WithAttribute(ctx, "key.generate.bits", "4096")

	public, private, err := generateKeyPair(4096)
	if err != nil {
		return "", "", err
	}

	return exportPubKeyAsPEM(public), exportPrivKeyAsPEM(private), nil
}

// Encrypt with public key.
func Encrypt(ctx context.Context, key, pass string) (string, error) {
	ctx = meta.WithAttribute(ctx, "key.encrypt.kind", "OAEP")
	meta.WithAttribute(ctx, "key.encrypt.hash", "sha512")

	block, _ := pem.Decode([]byte(key))

	k, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	e, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, k, []byte(pass), nil)
	if err != nil {
		return "", err
	}

	return string(e), nil
}

func generateKeyPair(bits int) (*rsa.PublicKey, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	return &privateKey.PublicKey, privateKey, nil
}

func exportPubKeyAsPEM(key *rsa.PublicKey) string {
	return string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(key),
		},
	))
}

func exportPrivKeyAsPEM(key *rsa.PrivateKey) string {
	return string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	))
}
