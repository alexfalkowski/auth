package grpc

import (
	"context"
	"encoding/base64"
	"strings"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/security/header"
	gmeta "github.com/alexfalkowski/go-service/transport/grpc/meta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GenerateAccessToken for gRPC.
func (s *Server) GenerateAccessToken(ctx context.Context, req *v1.GenerateAccessTokenRequest) (*v1.GenerateAccessTokenResponse, error) {
	resp := &v1.GenerateAccessTokenResponse{}

	id, p, err := s.idAndPassword(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	for _, a := range s.config.Admins {
		if a.ID == id && password.CompareHashAndPassword(ctx, a.Hash, p) == nil {
			p, h, err := s.passwordAndHash(ctx)
			if err != nil {
				return resp, status.Error(codes.Internal, err.Error())
			}

			b, err := key.Encrypt(ctx, s.config.Key.Public, p)
			if err != nil {
				return resp, status.Error(codes.Internal, err.Error())
			}

			resp.Token = &v1.AccessToken{Bearer: base64.StdEncoding.EncodeToString([]byte(b)), Password: &v1.Password{Plain: p, Hash: h}}
			resp.Meta = meta.Attributes(ctx)

			return resp, nil
		}
	}

	return nil, status.Error(codes.Unauthenticated, header.ErrInvalidAuthorization.Error())
}

func (s *Server) idAndPassword(ctx context.Context) (string, string, error) {
	md := gmeta.ExtractIncoming(ctx)

	values := md["authorization"]
	if len(values) == 0 {
		return "", "", header.ErrInvalidAuthorization
	}

	_, credentials, err := header.ParseAuthorization(values[0])
	if err != nil {
		return "", "", err
	}

	c, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		return "", "", err
	}

	t := strings.Split(string(c), ":")
	if len(t) != 2 {
		return "", "", header.ErrInvalidAuthorization
	}

	return t[0], t[1], nil
}
