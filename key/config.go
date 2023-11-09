package key

import (
	"os"
	"strings"
)

// Config for key.
type Config struct {
	RSA     Pair `yaml:"rsa" json:"rsa" toml:"rsa"`
	Ed25519 Pair `yaml:"ed25519" json:"ed25519" toml:"ed25519"`
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
	Public  string `yaml:"public"`
	Private string `yaml:"private"`
}

// GetPrivate from config or env
func (p Pair) GetPrivate() string {
	s := strings.Split(p.Private, ":")

	if len(s) != 2 || s[0] != "env" {
		return p.Private
	}

	return os.Getenv(s[1])
}
