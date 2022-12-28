package config

import (
	"strings"

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
	return &CasbinAdapter{policy: cfg.Casbin.Policy}
}

// NewCasbinEnforcer for config.
func NewCasbinEnforcer(m model.Model, a persist.Adapter) (*casbin.Enforcer, error) {
	return casbin.NewEnforcer(m, a)
}

// CasbinAdapter for casbin.
type CasbinAdapter struct {
	policy string
}

func (a *CasbinAdapter) AddPolicy(sec, ptype string, rule []string) error {
	return nil
}

func (a *CasbinAdapter) LoadPolicy(model model.Model) error {
	for _, p := range strings.Split(a.policy, "\n") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		if err := persist.LoadPolicyLine(p, model); err != nil {
			return err
		}
	}

	return nil
}

func (a *CasbinAdapter) RemoveFilteredPolicy(sec, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}

func (a *CasbinAdapter) RemovePolicy(sec, ptype string, rule []string) error {
	return nil
}

func (a *CasbinAdapter) SavePolicy(model model.Model) error {
	return nil
}
