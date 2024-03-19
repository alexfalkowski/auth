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

// RSA cipher.
type RSA struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// NewRSA cipher.
func NewRSA(publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) *RSA {
	return &RSA{publicKey: publicKey, privateKey: privateKey}
}

// Encrypt with RSA OAEP.
func (r *RSA) Encrypt(ctx context.Context, msg string) (string, error) {
	ctx = meta.WithAttribute(ctx, "keyEncryptKind", meta.Value("OAEP"))
	meta.WithAttribute(ctx, "keyEncryptHash", meta.Value("SHA512"))

	e, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, r.publicKey, []byte(msg), nil)

	return base64.StdEncoding.EncodeToString(e), err
}

// Decrypt with RSA OAEP.
func (r *RSA) Decrypt(ctx context.Context, cipher string) (string, error) {
	ctx = meta.WithAttribute(ctx, "keyEncryptKind", meta.Value("OAEP"))
	meta.WithAttribute(ctx, "keyEncryptHash", meta.Value("SHA512"))

	d, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, r.privateKey, []byte(cipher), nil)

	return string(d), err
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

// NewRSAPublicKey from key.
func NewRSAPublicKey(cfg *Config) (*rsa.PublicKey, error) {
	k, err := base64.StdEncoding.DecodeString(cfg.RSA.Public)
	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PublicKey(k)
}

// NewRSAPrivateKey from key.
func NewRSAPrivateKey(cfg *Config) (*rsa.PrivateKey, error) {
	k, err := base64.StdEncoding.DecodeString(cfg.RSA.GetPrivate())
	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PrivateKey(k)
}
