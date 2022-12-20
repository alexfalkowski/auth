package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// Generate public and private key or error.
func Generate() (string, string, error) {
	public, private, err := generateKeyPair(4096)
	if err != nil {
		return "", "", err
	}

	return exportPubKeyAsPEM(public), exportPrivKeyAsPEM(private), nil
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
