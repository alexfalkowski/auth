package token

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewJWT),
	fx.Provide(NewKID),
	fx.Provide(NewPaseto),
	fx.Provide(NewToken),
)
