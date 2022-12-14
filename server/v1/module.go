package v1

import (
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/auth/server/v1/transport/grpc"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(grpc.NewServer),
		fx.Invoke(grpc.Register),
		config.Module,
	)
)
