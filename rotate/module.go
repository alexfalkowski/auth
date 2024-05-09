package rotate

import (
	"go.uber.org/fx"
)

// CommandModule for fx.
var CommandModule = fx.Options(
	fx.Invoke(RunCommand),
)
