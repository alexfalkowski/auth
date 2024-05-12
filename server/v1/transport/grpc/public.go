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

	switch kind {
	case "rsa":
		resp.Meta = s.meta(ctx)
		resp.Key = string(s.rsaConfig.Public)

		return resp, nil
	case "ed25519":
		resp.Meta = s.meta(ctx)
		resp.Key = string(s.ed25519Config.Public)

		return resp, nil
	default:
		return resp, status.Errorf(codes.NotFound, "%s public key not found", kind)
	}
}
