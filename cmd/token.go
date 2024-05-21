package cmd

import (
	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/cmd/token"
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/go-service/compressor"
	"github.com/alexfalkowski/go-service/crypto"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"go.uber.org/fx"
)

// TokenOptions for cmd.
var TokenOptions = []fx.Option{
	runtime.Module, feature.Module,
	compressor.Module, marshaller.Module, crypto.Module,
	telemetry.Module, metrics.Module,
	config.Module, client.Module, token.Module, Module,
}
