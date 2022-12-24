package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/go-service/meta"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GenerateKey for gRPC.
func (s *Server) GenerateKey(ctx context.Context, req *v1.GenerateKeyRequest) (*v1.GenerateKeyResponse, error) {
	kind := req.Meta["kind"]
	if kind == "" {
		kind = "rsa"
	}

	resp := &v1.GenerateKeyResponse{}

	public, private, err := key.GeneratePair(kind)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Key = &v1.Key{Public: public, Private: private}

	resp.Meta = meta.Attributes(ctx)
	resp.Meta["kind"] = kind
	maps.Copy(resp.Meta, req.Meta)

	return resp, nil
}
