package password

import (
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

// Generate secure password.
func Generate() (string, error) {
	return password.Generate(64, 10, 10, false, false)
}

// Hash the password using bcrypt.
func Hash(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}
