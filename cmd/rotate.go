package cmd

import (
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/rotate"
	"github.com/alexfalkowski/auth/token"
	"github.com/alexfalkowski/go-service/telemetry"
	"go.uber.org/fx"
)

// RotateOptions for cmd.
var RotateOptions = []fx.Option{
	fx.NopLogger,
	telemetry.Module,
	Module, config.Module,
	key.Module, token.Module,
	password.Module, rotate.CommandModule,
}
