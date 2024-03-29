package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

// Config for casbin.
type Config struct {
	Model  string `yaml:"model,omitempty" json:"model,omitempty" toml:"model,omitempty"`
	Policy string `yaml:"policy,omitempty" json:"policy,omitempty" toml:"policy,omitempty"`
}

// NewCasbinModel for config.
func NewModel(cfg *Config) (model.Model, error) {
	return model.NewModelFromString(cfg.Model)
}

// NewCasbinAdapter for config.
func NewAdapter(cfg *Config) persist.Adapter {
	return &adapter{policy: cfg.Policy}
}

// NewCasbinEnforcer for config.
func NewEnforcer(m model.Model, a persist.Adapter) (*casbin.Enforcer, error) {
	return casbin.NewEnforcer(m, a)
}
