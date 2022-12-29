package config

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewRSAPrivateKey),
	fx.Provide(NewRSAPublicKey),
	fx.Provide(NewEd25519PrivateKey),
	fx.Provide(NewEd25519PublicKey),
	fx.Provide(NewBranca),
	fx.Provide(NewCasbinModel),
	fx.Provide(NewCasbinAdapter),
	fx.Provide(NewCasbinEnforcer),
)
