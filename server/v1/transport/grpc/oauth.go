package grpc

import (
	"context"
	"slices"
	"time"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/meta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GenerateOAuthToken for gRPC.
func (s *Server) GenerateOAuthToken(ctx context.Context, req *v1.GenerateOAuthTokenRequest) (*v1.GenerateOAuthTokenResponse, error) {
	resp := &v1.GenerateOAuthTokenResponse{}
	id := req.GetClientId()

	i := slices.IndexFunc(s.config.Services, func(svc *config.Service) bool { return svc.ID == id })
	if i == -1 {
		return resp, status.Error(codes.Unauthenticated, "missing service")
	}

	svc := s.config.Services[i]

	if err := s.secure.Compare(ctx, svc.Hash, req.GetClientSecret()); err != nil {
		return resp, status.Error(codes.Unauthenticated, err.Error())
	}

	d, err := time.ParseDuration(svc.Duration)
	if err != nil {
		return resp, err
	}

	to, err := s.generate("jwt", svc.Name, req.GetAudience(), d)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Meta = meta.Attributes(ctx)
	resp.AccessToken = to
	resp.TokenType = "Bearer"

	return resp, nil
}

// GetJWKSets for gRPC.
func (s *Server) GetJWKSets(ctx context.Context, _ *v1.GetJWKSetsRequest) (*v1.GetJWKSetsResponse, error) {
	resp := &v1.GetJWKSetsResponse{
		Meta: meta.Attributes(ctx),
		Keys: []*v1.JWKSet{
			{
				Kid: string(s.kid),
				Kty: "EC",
				Use: "sig",
				X5C: []string{s.key.Ed25519.Public},
			},
		},
	}

	return resp, nil
}
