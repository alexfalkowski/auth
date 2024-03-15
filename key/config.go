package key

import (
	"github.com/alexfalkowski/go-service/os"
)

// Config for key.
type Config struct {
	RSA     Pair `yaml:"rsa,omitempty" json:"rsa,omitempty" toml:"rsa,omitempty"`
	Ed25519 Pair `yaml:"ed25519,omitempty" json:"ed25519,omitempty" toml:"ed25519,omitempty"`
}

// Pair from kind.
func (c *Config) Pair(kind string) *Pair {
	switch kind {
	case "rsa":
		return &c.RSA
	case "ed25519":
		return &c.Ed25519
	default:
		return nil
	}
}

// Pair for key.
type Pair struct {
	Public  string `yaml:"public,omitempty" json:"public,omitempty" toml:"public,omitempty"`
	Private string `yaml:"private,omitempty" json:"private,omitempty" toml:"private,omitempty"`
}

// GetPrivate from config or env.
func (p Pair) GetPrivate() string {
	return os.GetFromEnv(p.Private)
}
