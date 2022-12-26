package cmd

import (
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/server/health"
	v1 "github.com/alexfalkowski/auth/server/v1"
	"github.com/alexfalkowski/auth/service"
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, fx.Provide(NewVersion), config.Module, health.Module,
	logger.ZapModule, metrics.PrometheusModule, transport.Module,
	key.Module, service.Module, v1.Module,
}
