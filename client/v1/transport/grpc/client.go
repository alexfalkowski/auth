package grpc

import (
	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/auth/transport/grpc"
	"github.com/alexfalkowski/go-service/env"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Client    *v1c.Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	UserAgent env.UserAgent
}

// NewServiceClient for gRPC.
//
//nolint:nilnil
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	if !v1c.IsEnabled(params.Client) {
		return nil, nil
	}

	opts := grpc.ClientOpts{
		Lifecycle: params.Lifecycle,
		Client:    params.Client.Config,
		Logger:    params.Logger,
		Tracer:    params.Tracer,
		Meter:     params.Meter,
		UserAgent: params.UserAgent,
	}
	conn, err := grpc.NewClient(opts)

	return v1.NewServiceClient(conn), err
}
