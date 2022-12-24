package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
)

// PrivateRSA from key.
func PrivateEd25519(key string) (ed25519.PrivateKey, error) {
	return base64.StdEncoding.DecodeString(key)
}

func generateEd25519() (string, string, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(pub), base64.StdEncoding.EncodeToString(pri), nil
}
