package password

import (
	"errors"
)

var (
	// ErrInvalidLength for password.
	ErrInvalidLength = errors.New("invalid length")

	// DefaultLength for password.
	DefaultLength = Length(64)
)

// Length to generate password.
type Length uint32

// Valid length.
func (l Length) Valid() error {
	if l < 32 {
		return ErrInvalidLength
	}

	return nil
}
