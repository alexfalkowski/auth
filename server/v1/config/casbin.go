package config

import (
	cas "github.com/alexfalkowski/auth/casbin"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

// NewCasbinModel for config.
func NewCasbinModel(cfg *Config) (model.Model, error) {
	return model.NewModelFromString(cfg.Casbin.Model)
}

// NewCasbinAdapter for config.
func NewCasbinAdapter(cfg *Config) persist.Adapter {
	return cas.NewAdapter(cfg.Casbin.Policy)
}

// NewCasbinEnforcer for config.
func NewCasbinEnforcer(m model.Model, a persist.Adapter) (*casbin.Enforcer, error) {
	return casbin.NewEnforcer(m, a)
}
