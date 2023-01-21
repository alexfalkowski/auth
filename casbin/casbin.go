package casbin

import (
	"strings"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type adapter struct {
	policy string
}

func (a *adapter) AddPolicy(sec, ptype string, rule []string) error {
	return nil
}

func (a *adapter) LoadPolicy(model model.Model) error {
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

func (a *adapter) RemoveFilteredPolicy(sec, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}

func (a *adapter) RemovePolicy(sec, ptype string, rule []string) error {
	return nil
}

func (a *adapter) SavePolicy(model model.Model) error {
	return nil
}
