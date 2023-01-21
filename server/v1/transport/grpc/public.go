package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/go-service/meta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetPublicKey from kind.
func (s *Server) GetPublicKey(ctx context.Context, req *v1.GetPublicKeyRequest) (*v1.GetPublicKeyResponse, error) {
	resp := &v1.GetPublicKeyResponse{}

	pair := s.key.Pair(req.Kind)
	if pair == nil {
		return resp, status.Errorf(codes.NotFound, "%s public key not found", req.Kind)
	}

	resp.Meta = meta.Attributes(ctx)
	resp.Key = pair.Public

	return resp, nil
}
