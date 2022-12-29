package config

import (
	"github.com/essentialkaos/branca"
)

// NewBranca for config.
func NewBranca(cfg *Config) (*branca.Branca, error) {
	return branca.NewBranca([]byte(cfg.Secret.Branca))
}
