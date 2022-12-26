package grpc

import (
	"context"
	"encoding/base64"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/service"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/security/header"
	gmeta "github.com/alexfalkowski/go-service/transport/grpc/meta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GenerateServiceToken for gRPC.
func (s *Server) GenerateServiceToken(ctx context.Context, req *v1.GenerateServiceTokenRequest) (*v1.GenerateServiceTokenResponse, error) {
	kind := req.Kind
	if kind == "" {
		kind = "jwt"
	}

	resp := &v1.GenerateServiceTokenResponse{}

	p, err := s.password(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	for _, svc := range s.config.Services {
		if password.CompareHashAndPassword(ctx, svc.Hash, p) == nil {
			params := service.TokenParams{
				RSA:     service.KeyPair{Public: s.config.Key.RSA.Public, Private: s.config.Key.RSA.Private},
				Ed25519: service.KeyPair{Public: s.config.Key.Ed25519.Public, Private: s.config.Key.Ed25519.Private},
				Branca:  s.config.Secret.Branca,
				Service: service.Service{ID: svc.ID, Hash: svc.Hash, Duration: svc.Duration},
				Issuer:  s.config.Issuer,
				Kind:    kind,
			}

			to, err := service.GenerateToken(params)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

			resp.Meta = meta.Attributes(ctx)
			resp.Meta["kind"] = kind

			resp.Token = &v1.ServiceToken{Bearer: to}

			return resp, nil
		}
	}

	return nil, status.Error(codes.Unauthenticated, header.ErrInvalidAuthorization.Error())
}

func (s *Server) password(ctx context.Context) (string, error) {
	md := gmeta.ExtractIncoming(ctx)

	values := md["authorization"]
	if len(values) == 0 {
		return "", header.ErrInvalidAuthorization
	}

	_, credentials, err := header.ParseAuthorization(values[0])
	if err != nil {
		return "", err
	}

	c, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		return "", err
	}

	return key.DecryptRSA(ctx, s.config.Key.RSA.Private, string(c))
}
