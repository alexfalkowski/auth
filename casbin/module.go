package casbin

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewAdapter),
	fx.Provide(NewEnforcer),
	fx.Provide(NewModel),
)
