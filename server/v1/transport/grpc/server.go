package grpc

import (
	"context"
	"encoding/base64"
	"strings"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	sv1 "github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/security/header"
	"github.com/alexfalkowski/go-service/transport/grpc/meta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewServer for gRPC.
func NewServer(config *sv1.Config) v1.ServiceServer {
	return &Server{config: config}
}

// Server for gRPC.
type Server struct {
	config *sv1.Config
	v1.UnimplementedServiceServer
}

// GeneratePassword for gRPC.
func (s *Server) GeneratePassword(ctx context.Context, req *v1.GeneratePasswordRequest) (*v1.GeneratePasswordResponse, error) {
	resp := &v1.GeneratePasswordResponse{}

	p, h, err := s.passwordAndHash()
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Password = &v1.Password{Plain: p, Hash: h}

	return resp, nil
}

// GenerateKey for gRPC.
func (s *Server) GenerateKey(ctx context.Context, req *v1.GenerateKeyRequest) (*v1.GenerateKeyResponse, error) {
	resp := &v1.GenerateKeyResponse{}

	public, private, err := key.Generate()
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Key = &v1.Key{Public: public, Private: private}

	return resp, nil
}

// GenerateAccessToken for gRPC.
func (s *Server) GenerateAccessToken(ctx context.Context, req *v1.GenerateAccessTokenRequest) (*v1.GenerateAccessTokenResponse, error) {
	resp := &v1.GenerateAccessTokenResponse{}

	id, p, err := s.idAndPassword(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	for _, a := range s.config.Admins {
		if a.ID == id && password.CompareHashAndPassword(a.Hash, p) == nil {
			p, h, err := s.passwordAndHash()
			if err != nil {
				return resp, status.Error(codes.Internal, err.Error())
			}

			b, err := key.Encrypt(s.config.Key.Public, p)
			if err != nil {
				return resp, status.Error(codes.Internal, err.Error())
			}

			resp.Token = &v1.AccessToken{Bearer: base64.StdEncoding.EncodeToString([]byte(b)), Password: &v1.Password{Plain: p, Hash: h}}

			return resp, nil
		}
	}

	return nil, status.Error(codes.Unauthenticated, header.ErrInvalidAuthorization.Error())
}

func (s *Server) passwordAndHash() (string, string, error) {
	p, err := password.Generate()
	if err != nil {
		return "", "", err
	}

	h, err := password.Hash(p)
	if err != nil {
		return "", "", err
	}

	return p, h, nil
}

func (s *Server) idAndPassword(ctx context.Context) (string, string, error) {
	md := meta.ExtractIncoming(ctx)

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
