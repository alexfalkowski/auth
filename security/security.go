package security

import (
	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/go-service/security"
	"github.com/alexfalkowski/go-service/security/token"
)

// NewToken for security.
func NewToken(c *token.Config, cli *client.Client) (token.Generator, token.Verifier) {
	kind := "auth"

	token.RegisterGenerator(kind, client.NewGenerator(cli))
	token.RegisterVerifier(kind, client.NewVerifier(cli))

	return security.NewToken(c)
}
