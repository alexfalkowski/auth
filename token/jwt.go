package token

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// KID is a key ID.
type KID string

// NewKID for JWKSets.
func NewKID() (KID, error) {
	bytes := make([]byte, 10)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return KID(hex.EncodeToString(bytes)), nil
}

// JWT token.
type JWT struct {
	kid        KID
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

// NewJWT token.
func NewJWT(kid KID, publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey) *JWT {
	return &JWT{kid: kid, publicKey: publicKey, privateKey: privateKey}
}

// Generate JWT token.
func (j *JWT) Generate(sub, aud, iss string, exp time.Duration) (string, error) {
	t := time.Now()

	claims := &jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: t.Add(exp)},
		ID:        uuid.NewString(),
		IssuedAt:  &jwt.NumericDate{Time: t},
		Issuer:    iss,
		NotBefore: &jwt.NumericDate{Time: t},
		Subject:   sub,
		Audience:  []string{aud},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	token.Header["kid"] = j.kid

	return token.SignedString(j.privateKey)
}

// Verify JWT token.
func (j *JWT) Verify(token, aud, iss string) (string, error) {
	claims := &jwt.RegisteredClaims{}

	t, err := jwt.ParseWithClaims(token, claims, j.validate)
	if err != nil {
		return "", err
	}

	if t.Header["alg"] != "EdDSA" {
		return "", ErrInvalidAlgorithm
	}

	if !claims.VerifyIssuer(iss, true) {
		return "", ErrInvalidIssuer
	}

	if !claims.VerifyAudience(aud, true) {
		return "", ErrInvalidAudience
	}

	return claims.Subject, nil
}

func (j *JWT) validate(_ *jwt.Token) (any, error) {
	return j.publicKey, nil
}
