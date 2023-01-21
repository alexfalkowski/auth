package key

import (
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewRSA),
	fx.Provide(NewEd25519),
	fx.Provide(NewGenerator),
	fx.Provide(NewRSAPrivateKey),
	fx.Provide(NewRSAPublicKey),
	fx.Provide(NewEd25519PrivateKey),
	fx.Provide(NewEd25519PublicKey),
)
