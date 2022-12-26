package config

import (
	"github.com/alexfalkowski/auth/service"
)

// NewBranca for config.
func NewBranca(cfg *Config) service.BrancaSecret {
	return service.BrancaSecret(cfg.Secret.Branca)
}
