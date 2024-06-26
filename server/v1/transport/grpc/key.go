package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GenerateKey for gRPC.
func (s *Server) GenerateKey(ctx context.Context, req *v1.GenerateKeyRequest) (*v1.GenerateKeyResponse, error) {
	kind := req.GetKind()
	if kind == "" {
		kind = "rsa"
	}

	resp := &v1.GenerateKeyResponse{}

	public, private, err := s.key.Generate(kind)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Key = &v1.Key{Public: public, Private: private}

	resp.Meta = s.meta(ctx)
	resp.Meta["kind"] = kind

	return resp, nil
}
