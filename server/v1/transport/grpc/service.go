package grpc

import (
	"context"
	"encoding/base64"
	"fmt"
	"slices"
	"time"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/server/v1/config"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/transport/grpc/security/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GenerateServiceToken for gRPC.
func (s *Server) GenerateServiceToken(ctx context.Context, req *v1.GenerateServiceTokenRequest) (*v1.GenerateServiceTokenResponse, error) {
	kind := req.GetKind()
	if kind == "" {
		kind = "jwt"
	}

	resp := &v1.GenerateServiceTokenResponse{}

	p, err := s.password(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	i := slices.IndexFunc(s.config.Services, func(svc *config.Service) bool { return s.secure.Compare(ctx, svc.Hash, p) == nil })
	if i == -1 {
		return resp, status.Error(codes.Unauthenticated, "missing service")
	}

	svc := s.config.Services[i]

	d, err := time.ParseDuration(svc.Duration)
	if err != nil {
		return resp, err
	}

	to, err := s.generate(kind, svc.Name, req.GetAudience(), d)
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}

	resp.Meta = meta.Attributes(ctx)
	resp.Meta["kind"] = kind

	resp.Token = &v1.ServiceToken{Bearer: to}

	return resp, nil
}

// VerifyServiceToken for gRPC.
func (s *Server) VerifyServiceToken(ctx context.Context, req *v1.VerifyServiceTokenRequest) (*v1.VerifyServiceTokenResponse, error) {
	kind := req.GetKind()
	if kind == "" {
		kind = "jwt"
	}

	resp := &v1.VerifyServiceTokenResponse{}
	aud := req.GetAudience()
	act := req.GetAction()

	t, err := token.ExtractToken(ctx)
	if err != nil {
		return resp, status.Error(codes.Unauthenticated, err.Error())
	}

	sub, err := s.svc.Verify(t, kind, aud, s.config.Issuer)
	if err != nil {
		return resp, status.Error(codes.Unauthenticated, err.Error())
	}

	ok, err := s.enforcer.Enforce(sub, aud, act)
	if err != nil {
		return resp, status.Error(codes.Unauthenticated, err.Error())
	}

	if !ok {
		return resp, status.Errorf(codes.Unauthenticated, "enforcing %s %s %s failed", sub, aud, act)
	}

	resp.Meta = meta.Attributes(ctx)
	resp.Meta["kind"] = kind

	return resp, nil
}

func (s *Server) password(ctx context.Context) (string, error) {
	credentials, err := token.ExtractToken(ctx)
	if err != nil {
		return "", err
	}

	c, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		return "", err
	}

	return s.rsa.Decrypt(ctx, string(c))
}

func (s *Server) generate(kind, sub, aud string, exp time.Duration) (string, error) {
	key := fmt.Sprintf("%s:%s:%s", kind, sub, aud)

	v, ok := s.cache.Get(key)
	if ok {
		return v.(string), nil
	}

	t, err := s.svc.Generate(kind, sub, aud, s.config.Issuer, exp)
	if err != nil {
		return "", err
	}

	s.cache.SetWithTTL(key, t, 0, exp)

	return t, nil
}
