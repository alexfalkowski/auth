package cmd

import (
	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	fx.NopLogger, runtime.Module, Module, telemetry.Module, metrics.Module,
	config.Module, transport.Module, client.CommandModule,
}
