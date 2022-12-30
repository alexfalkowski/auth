package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/go-service/meta"
)

// GeneratePassword for gRPC.
func (s *Server) GeneratePassword(ctx context.Context, req *v1.GeneratePasswordRequest) (*v1.GeneratePasswordResponse, error) {
	length := req.Length
	if length == 0 {
		length = uint32(password.DefaultLength)
	}

	resp := &v1.GeneratePasswordResponse{}

	p, h, err := s.passwordAndHash(ctx, length)
	if err != nil {
		return resp, err
	}

	resp.Password = &v1.Password{Plain: p, Hash: h}
	resp.Meta = meta.Attributes(ctx)

	return resp, nil
}
