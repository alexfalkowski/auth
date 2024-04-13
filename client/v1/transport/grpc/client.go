package grpc

import (
	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/auth/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	ClientConfig *v1c.Config
	Logger       *zap.Logger
	Tracer       trace.Tracer
	Meter        metric.Meter
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	cfg := params.ClientConfig
	if cfg == nil {
		return nil, nil
	}

	opts := grpc.ClientOpts{
		Lifecycle:    params.Lifecycle,
		ClientConfig: cfg.Config,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
	}

	conn, err := grpc.NewClient(opts)
	if err != nil {
		return nil, err
	}

	return v1.NewServiceClient(conn), nil
}
