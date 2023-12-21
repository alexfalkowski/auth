package cmd

import (
	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/transport"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	fx.NopLogger, runtime.Module, Module, debug.Module, feature.Module,
	telemetry.Module, metrics.Module,
	config.Module, transport.Module, client.CommandModule,
}
