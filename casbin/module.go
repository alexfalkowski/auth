package casbin

import (
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(NewAdapter),
		fx.Provide(NewEnforcer),
		fx.Provide(NewModel),
	)
)
