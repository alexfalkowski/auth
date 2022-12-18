package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/password"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewServer for gRPC.
func NewServer() v1.ServiceServer {
	return &Server{}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
}

// GeneratePassword for gRPC.
func (s *Server) GeneratePassword(ctx context.Context, req *v1.GeneratePasswordRequest) (*v1.GeneratePasswordResponse, error) {
	resp := &v1.GeneratePasswordResponse{}

	p, err := password.Generate()
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	h, err := password.Hash(p)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Password = &v1.Password{Plain: p, Hash: h}

	return resp, nil
}
