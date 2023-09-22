package cmd

import (
	"github.com/alexfalkowski/auth/casbin"
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/health"
	v1 "github.com/alexfalkowski/auth/server/v1"
	"github.com/alexfalkowski/auth/service"
	"github.com/alexfalkowski/go-service/cache"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, runtime.Module, telemetry.Module, metrics.Module, marshaller.Module,
	cache.RistrettoModule, Module, config.Module, health.Module,
	transport.Module, key.Module, casbin.Module, service.Module,
	password.Module, v1.Module,
}
