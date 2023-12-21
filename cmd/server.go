package cmd

import (
	"github.com/alexfalkowski/auth/casbin"
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/health"
	v1 "github.com/alexfalkowski/auth/server/v1"
	"github.com/alexfalkowski/auth/token"
	"github.com/alexfalkowski/auth/transport"
	"github.com/alexfalkowski/go-service/cache"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, runtime.Module, debug.Module, feature.Module,
	telemetry.Module, metrics.Module, health.Module,
	cache.RistrettoModule, Module, config.Module,
	transport.Module, key.Module, casbin.Module, token.Module,
	password.Module, v1.Module,
}
