package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetPublicKey from kind.
func (s *Server) GetPublicKey(ctx context.Context, req *v1.GetPublicKeyRequest) (*v1.GetPublicKeyResponse, error) {
	resp := &v1.GetPublicKeyResponse{}
	kind := req.GetKind()

	pair := s.keyConfig.Pair(kind)
	if pair == nil {
		return resp, status.Errorf(codes.NotFound, "%s public key not found", kind)
	}

	resp.Meta = s.meta(ctx)
	resp.Key = pair.Public

	return resp, nil
}
