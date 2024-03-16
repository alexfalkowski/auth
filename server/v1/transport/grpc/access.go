package grpc

import (
	"context"
	"encoding/base64"
	"slices"
	"strings"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/security/header"
	"github.com/alexfalkowski/go-service/transport/grpc/security/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GenerateAccessToken for gRPC.
func (s *Server) GenerateAccessToken(ctx context.Context, req *v1.GenerateAccessTokenRequest) (*v1.GenerateAccessTokenResponse, error) {
	length := req.GetLength()
	if length == 0 {
		length = uint32(password.DefaultLength)
	}

	resp := &v1.GenerateAccessTokenResponse{}

	id, p, err := s.idAndPassword(ctx)
	if err != nil {
		return resp, status.Error(codes.Unauthenticated, err.Error())
	}

	i := slices.IndexFunc(s.config.Admins, func(a *config.Admin) bool { return a.ID == id })
	if i == -1 {
		return resp, status.Error(codes.Unauthenticated, "missing admin")
	}

	a := s.config.Admins[i]

	if err := s.secure.Compare(ctx, a.Hash, p); err != nil {
		return resp, status.Error(codes.Unauthenticated, err.Error())
	}

	p, h, err := s.passwordAndHash(ctx, length)
	if err != nil {
		return resp, err
	}

	b, err := s.rsa.Encrypt(ctx, p)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Token = &v1.AccessToken{Bearer: b, Password: &v1.Password{Plain: p, Hash: h}}
	resp.Meta = meta.Attributes(ctx)

	return resp, nil
}

func (s *Server) idAndPassword(ctx context.Context) (string, string, error) {
	credentials, err := token.ExtractToken(ctx)
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
