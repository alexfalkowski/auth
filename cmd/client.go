package cmd

import (
	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/otel"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	fx.NopLogger, runtime.Module, Module, marshaller.Module, otel.Module,
	config.Module, logger.ZapModule, transport.GRPCModule, client.CommandModule,
}
