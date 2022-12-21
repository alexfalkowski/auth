package config

import (
	"github.com/alexfalkowski/go-service/config"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(NewConfigurator),
		config.UnmarshalModule,
		config.ConfigModule,
		config.WatchModule,
		fx.Provide(v1Config),
		fx.Provide(healthConfig),
	)
)
