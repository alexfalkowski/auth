package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

// NewRSAPublicKey from key.
func NewRSAPublicKey(cfg *Config) (*rsa.PublicKey, error) {
	k, err := base64.StdEncoding.DecodeString(cfg.Key.RSA.Public)
	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PublicKey(k)
}

// NewRSAPrivateKey from key.
func NewRSAPrivateKey(cfg *Config) (*rsa.PrivateKey, error) {
	k, err := base64.StdEncoding.DecodeString(cfg.Key.RSA.Private)
	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PrivateKey(k)
}
