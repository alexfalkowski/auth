package grpc

import (
	"context"
	"errors"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/go-service/meta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		if errors.Is(err, password.ErrInvalidLength) {
			return resp, status.Error(codes.InvalidArgument, err.Error())
		}

		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Password = &v1.Password{Plain: p, Hash: h}
	resp.Meta = meta.Attributes(ctx)

	return resp, nil
}
