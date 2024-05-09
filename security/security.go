package security

import (
	"github.com/alexfalkowski/go-service/security/token"
)

// IsAuth security.
func IsAuth(c *token.Config) bool {
	return c != nil && c.Kind == "auth"
}
