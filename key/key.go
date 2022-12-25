package key

import (
	"errors"
)

var (
	// ErrInvalidKind for key.
	ErrInvalidKind = errors.New("invalid kind")
)

// GeneratePair for key.
func GeneratePair(kind string) (string, string, error) {
	switch kind {
	case "rsa":
		return generateRSA()
	case "ed25519":
		return generateEd25519()
	}

	return "", "", ErrInvalidKind
}
