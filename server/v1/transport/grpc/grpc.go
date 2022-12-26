package grpc

import (
	"context"
	"crypto/ed25519"
	"fmt"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	sv1 "github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/auth/service"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	"github.com/alexfalkowski/go-service/transport/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle        fx.Lifecycle
	GRPCServer       *grpc.Server
	HTTPServer       *http.Server
	GRPCConfig       *grpc.Config
	TransportConfig  *transport.Config
	Logger           *zap.Logger
	Tracer           opentracing.Tracer
	Metrics          *prometheus.ClientMetrics
	V1Config         *sv1.Config
	RSA              *key.RSA
	KeyGenerator     *key.Generator
	ServiceGenerator *service.Generator
	PrivateKey       ed25519.PrivateKey
}

// Register server.
func Register(params RegisterParams) error {
	ctx := context.Background()
	server := NewServer(params.V1Config, params.RSA, params.KeyGenerator, params.ServiceGenerator, params.PrivateKey)

	v1.RegisterServiceServer(params.GRPCServer.Server, server)

	conn, err := grpc.NewClient(
		grpc.ClientParams{Context: ctx, Host: fmt.Sprintf("127.0.0.1:%s", params.TransportConfig.Port), Config: params.GRPCConfig},
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer), grpc.WithClientMetrics(params.Metrics),
	)
	if err != nil {
		return err
	}

	if err := v1.RegisterServiceHandler(ctx, params.HTTPServer.Mux, conn); err != nil {
		return err
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

	return nil
}
