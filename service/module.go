package service

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewJWT),
	fx.Provide(NewPaseto),
	fx.Provide(NewService),
)
