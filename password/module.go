package password

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewGeneratorInput),
	fx.Provide(NewGenerator),
	fx.Provide(NewSecure),
)
