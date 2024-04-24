package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	g "github.com/alexfalkowski/auth/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Mux       *runtime.ServeMux
	GRPC      *grpc.Server
	Client    *v1c.Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	Server    v1.ServiceServer
}

// Register server.
func Register(params RegisterParams) error {
	v1.RegisterServiceServer(params.GRPC.Server(), params.Server)

	ctx := context.Background()
	opts := g.ClientOpts{
		Lifecycle: params.Lifecycle,
		Client:    params.Client.Config,
		Logger:    params.Logger,
		Tracer:    params.Tracer,
		Meter:     params.Meter,
	}

	conn, err := g.NewClient(opts)
	if err != nil {
		return err
	}

	return v1.RegisterServiceHandler(ctx, params.Mux, conn)
}
