package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *grpc.Config
	Logger    *zap.Logger
	Tracer    tracer.Tracer
	Client    *v1c.Config
	Meter     metric.Meter
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	opts := []grpc.ClientOption{
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer),
		grpc.WithClientMetrics(params.Meter), grpc.WithClientRetry(),
	}

	if params.Config.Security.IsEnabled() {
		sec, err := grpc.WithClientSecure(params.Config.Security)
		if err != nil {
			return nil, err
		}

		opts = append(opts, sec)
	}

	conn, err := grpc.NewClient(context.Background(), params.Client.Host, params.Config, opts...)
	if err != nil {
		return nil, err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn.ResetConnectBackoff()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})

	return v1.NewServiceClient(conn), nil
}
