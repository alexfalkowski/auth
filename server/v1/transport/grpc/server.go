package grpc

import (
	"context"
	"errors"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/auth/token"
	"github.com/alexfalkowski/go-service/crypto/ed25519"
	"github.com/alexfalkowski/go-service/crypto/rsa"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/ristretto"
	"github.com/casbin/casbin/v2"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerParams for gRPC.
type ServerParams struct {
	fx.In

	Ed25519Config *ed25519.Config
	RSAConfig     *rsa.Config
	Config        *config.Config
	RSA           *key.RSA
	Key           *key.Generator
	Token         *token.Token
	KID           token.KID
	Secure        *password.Secure
	Enforcer      *casbin.Enforcer
	Cache         ristretto.Cache
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{
		config: params.Config, ed25519Config: params.Ed25519Config, rsaConfig: params.RSAConfig,
		rsa: params.RSA, key: params.Key, token: params.Token, kid: params.KID, secure: params.Secure,
		enforcer: params.Enforcer, cache: params.Cache,
	}
}

// Server for gRPC.
type Server struct {
	ed25519Config *ed25519.Config
	rsaConfig     *rsa.Config
	config        *config.Config
	rsa           *key.RSA
	key           *key.Generator
	token         *token.Token
	kid           token.KID
	secure        *password.Secure
	enforcer      *casbin.Enforcer
	cache         ristretto.Cache

	v1.UnimplementedServiceServer
}

func (s *Server) passwordAndHash(length uint32) (string, string, error) {
	p, err := s.secure.Generate(password.Length(length))
	if err != nil {
		if errors.Is(err, password.ErrInvalidLength) {
			return "", "", status.Error(codes.InvalidArgument, err.Error())
		}

		return "", "", status.Error(codes.Internal, err.Error())
	}

	h, err := s.secure.Hash(p)
	if err != nil {
		return "", "", status.Error(codes.Internal, err.Error())
	}

	return p, h, nil
}

func (s *Server) meta(ctx context.Context) map[string]string {
	return meta.CamelStrings(ctx, "")
}
