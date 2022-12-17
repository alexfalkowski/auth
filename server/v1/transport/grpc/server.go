package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
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

	p, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Password = &v1.Password{Plain: p, Hash: string(h)}

	return resp, nil
}
