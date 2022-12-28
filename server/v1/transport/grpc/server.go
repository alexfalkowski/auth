package grpc

import (
	"context"
	"crypto/ed25519"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/auth/service"
	"github.com/alexfalkowski/go-service/security/header"
	"github.com/alexfalkowski/go-service/transport/grpc/meta"
	"github.com/casbin/casbin/v2"
	"go.uber.org/fx"
)

// ServerParams for gRPC.
type ServerParams struct {
	fx.In

	Config           *config.Config
	RSA              *key.RSA
	KeyGenerator     *key.Generator
	ServiceGenerator *service.Service
	PrivateKey       ed25519.PrivateKey
	Secure           *password.Secure
	Enforcer         *casbin.Enforcer
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{
		config: params.Config, rsa: params.RSA, key: params.KeyGenerator,
		service: params.ServiceGenerator, privateKey: params.PrivateKey,
		enforcer: params.Enforcer,
	}
}

// Server for gRPC.
type Server struct {
	config     *config.Config
	rsa        *key.RSA
	key        *key.Generator
	service    *service.Service
	privateKey ed25519.PrivateKey
	secure     *password.Secure
	enforcer   *casbin.Enforcer

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

func (s *Server) credentials(ctx context.Context) (string, error) {
	md := meta.ExtractIncoming(ctx)

	values := md["authorization"]
	if len(values) == 0 {
		return "", header.ErrInvalidAuthorization
	}

	_, credentials, err := header.ParseAuthorization(values[0])
	if err != nil {
		return "", err
	}

	return credentials, nil
}
