package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/security"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// ClientOpts for gRPC.
type ClientOpts struct {
	Lifecycle    fx.Lifecycle
	ClientConfig *client.Config
	Logger       *zap.Logger
	Tracer       trace.Tracer
	Meter        metric.Meter
}

// NewClient for gRPC.
func NewClient(options ClientOpts) (*g.ClientConn, error) {
	cfg := options.ClientConfig
	opts := []grpc.ClientOption{
		grpc.WithClientLogger(options.Logger), grpc.WithClientTracer(options.Tracer),
		grpc.WithClientMetrics(options.Meter), grpc.WithClientRetry(cfg.Retry),
		grpc.WithClientUserAgent(cfg.UserAgent),
	}

	if security.IsEnabled(cfg.Security) {
		sec, err := grpc.WithClientSecure(cfg.Security)
		if err != nil {
			return nil, err
		}

		opts = append(opts, sec)
	}

	conn, err := grpc.NewClient(cfg.Host, opts...)
	if err != nil {
		return nil, err
	}

	options.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			conn.ResetConnectBackoff()

			return nil
		},
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, nil
}
