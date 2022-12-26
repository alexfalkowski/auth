package grpc

import (
	"context"
	"crypto/ed25519"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/auth/service"
)

// NewServer for gRPC.
func NewServer(config *config.Config, rsa *key.RSA, kgen *key.Generator, sgen *service.Generator, pkey ed25519.PrivateKey) v1.ServiceServer {
	return &Server{config: config, rsa: rsa, kgen: kgen, sgen: sgen, privateKey: pkey}
}

// Server for gRPC.
type Server struct {
	config     *config.Config
	rsa        *key.RSA
	kgen       *key.Generator
	sgen       *service.Generator
	privateKey ed25519.PrivateKey
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
