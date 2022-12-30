package password

import (
	"crypto/rand"

	"github.com/sethvargo/go-password/password"
)

// NewGeneratorInput for password.
func NewGeneratorInput() *password.GeneratorInput {
	return &password.GeneratorInput{
		Reader: rand.Reader,
	}
}

// NewGenerator for password.
func NewGenerator(input *password.GeneratorInput) (*password.Generator, error) {
	return password.NewGenerator(input)
}
