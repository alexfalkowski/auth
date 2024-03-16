package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/go-service/security/token"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	ClientConfig *v1c.Config
	TokenConfig  *token.Config
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	if params.TokenConfig.Kind != "auth" {
		return nil, nil
	}

	opts := []grpc.ClientOption{
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer),
		grpc.WithClientMetrics(params.Meter), grpc.WithClientRetry(params.ClientConfig.Retry),
		grpc.WithClientUserAgent(params.ClientConfig.UserAgent),
	}

	if params.ClientConfig.Security.Enabled {
		sec, err := grpc.WithClientSecure(params.ClientConfig.Security)
		if err != nil {
			return nil, err
		}

		opts = append(opts, sec)
	}

	conn, err := grpc.NewClient(context.Background(), params.ClientConfig.Host, opts...)
	if err != nil {
		return nil, err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			conn.ResetConnectBackoff()

			return nil
		},
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return v1.NewServiceClient(conn), nil
}
