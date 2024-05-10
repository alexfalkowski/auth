package token

import (
	"encoding/hex"
	"time"

	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// KID is a key ID.
type KID string

// NewKID for JWKSets.
func NewKID() (KID, error) {
	b, err := rand.GenerateBytes(10)
	if err != nil {
		return "", err
	}

	return KID(hex.EncodeToString(b)), nil
}

// JWT token.
type JWT struct {
	ed  *key.Ed25519
	kid KID
}

// NewJWT token.
func NewJWT(kid KID, ed *key.Ed25519) *JWT {
	return &JWT{kid: kid, ed: ed}
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

	return token.SignedString(j.ed.Algo().PrivateKey())
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
	return j.ed.Algo().PublicKey(), nil
}
