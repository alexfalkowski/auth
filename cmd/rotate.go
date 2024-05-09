package cmd

import (
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/rotate"
	"github.com/alexfalkowski/auth/token"
	"github.com/alexfalkowski/go-service/compressor"
	"github.com/alexfalkowski/go-service/crypto"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"go.uber.org/fx"
)

// RotateOptions for cmd.
var RotateOptions = []fx.Option{
	runtime.Module, feature.Module,
	telemetry.Module, metrics.Module,
	compressor.Module, marshaller.Module, crypto.Module,
	config.Module, key.Module, token.Module,
	password.Module, rotate.CommandModule, Module,
}
