package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/v1/config"
)

// NewServer for gRPC.
func NewServer(config *config.Config) v1.ServiceServer {
	return &Server{config: config}
}

// Server for gRPC.
type Server struct {
	config *config.Config
	v1.UnimplementedServiceServer
}

func (s *Server) passwordAndHash(ctx context.Context) (string, string, error) {
	p, err := password.Generate(ctx)
	if err != nil {
		return "", "", err
	}

	h, err := password.Hash(ctx, p)
	if err != nil {
		return "", "", err
	}

	return p, h, nil
}
