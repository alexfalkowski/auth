package grpc

import (
	"context"
	"crypto/ed25519"
	"errors"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/auth/service"
	"github.com/casbin/casbin/v2"
	"github.com/dgraph-io/ristretto"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerParams for gRPC.
type ServerParams struct {
	fx.In

	Config           *config.Config
	Key              *key.Config
	RSA              *key.RSA
	KeyGenerator     *key.Generator
	ServiceGenerator *service.Service
	KID              service.KID
	PrivateKey       ed25519.PrivateKey
	Secure           *password.Secure
	Enforcer         *casbin.Enforcer
	Cache            *ristretto.Cache
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{
		config: params.Config, key: params.Key,
		rsa: params.RSA, gen: params.KeyGenerator,
		svc: params.ServiceGenerator, pvk: params.PrivateKey, kid: params.KID,
		enforcer: params.Enforcer, cache: params.Cache,
	}
}

// Server for gRPC.
type Server struct {
	config   *config.Config
	key      *key.Config
	rsa      *key.RSA
	gen      *key.Generator
	svc      *service.Service
	kid      service.KID
	pvk      ed25519.PrivateKey
	secure   *password.Secure
	enforcer *casbin.Enforcer
	cache    *ristretto.Cache

	v1.UnimplementedServiceServer
}

func (s *Server) passwordAndHash(ctx context.Context, length uint32) (string, string, error) {
	p, err := s.secure.Generate(ctx, password.Length(length))
	if err != nil {
		if errors.Is(err, password.ErrInvalidLength) {
			return "", "", status.Error(codes.InvalidArgument, err.Error())
		}

		return "", "", status.Error(codes.Internal, err.Error())
	}

	h, err := s.secure.Hash(ctx, p)
	if err != nil {
		return "", "", status.Error(codes.Internal, err.Error())
	}

	return p, h, nil
}
