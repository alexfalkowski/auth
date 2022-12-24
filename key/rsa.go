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

// PrivateRSA from key.
func PrivateRSA(key string) (*rsa.PrivateKey, error) {
	k, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PrivateKey(k)
}

// EncryptRSA with public key.
func EncryptRSA(ctx context.Context, key, pass string) (string, error) {
	ctx = meta.WithAttribute(ctx, "key.encrypt.kind", "OAEP")
	meta.WithAttribute(ctx, "key.encrypt.hash", "SHA512")

	k, err := publicRSA(key)
	if err != nil {
		return "", err
	}

	e, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, k, []byte(pass), nil)
	if err != nil {
		return "", err
	}

	return string(e), nil
}

// DecryptRSA with private key.
func DecryptRSA(ctx context.Context, key, cipher string) (string, error) {
	ctx = meta.WithAttribute(ctx, "key.decrypt.kind", "OAEP")
	meta.WithAttribute(ctx, "key.decrypt.hash", "SHA512")

	k, err := PrivateRSA(key)
	if err != nil {
		return "", err
	}

	d, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, k, []byte(cipher), nil)
	if err != nil {
		return "", err
	}

	return string(d), nil
}

func generateRSA() (string, string, error) {
	p, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}

	pub := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&p.PublicKey))
	pri := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(p))

	return pub, pri, nil
}

func publicRSA(key string) (*rsa.PublicKey, error) {
	k, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PublicKey(k)
}
