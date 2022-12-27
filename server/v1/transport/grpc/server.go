package grpc

import (
	"context"
	"crypto/ed25519"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/auth/service"
	"go.uber.org/fx"
)

// ServerParams for gRPC.
type ServerParams struct {
	fx.In

	Config           *config.Config
	RSA              *key.RSA
	KeyGenerator     *key.Generator
	ServiceGenerator *service.Generator
	PrivateKey       ed25519.PrivateKey
	Secure           *password.Secure
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{
		config: params.Config, rsa: params.RSA, key: params.KeyGenerator,
		service: params.ServiceGenerator, privateKey: params.PrivateKey,
	}
}

// Server for gRPC.
type Server struct {
	config     *config.Config
	rsa        *key.RSA
	key        *key.Generator
	service    *service.Generator
	privateKey ed25519.PrivateKey
	secure     *password.Secure

	v1.UnimplementedServiceServer
}

func (s *Server) passwordAndHash(ctx context.Context) (string, string, error) {
	p, err := s.secure.Generate(ctx)
	if err != nil {
		return "", "", err
	}

	h, err := s.secure.Hash(ctx, p)
	if err != nil {
		return "", "", err
	}

	return p, h, nil
}
