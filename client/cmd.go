package client

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

var (
	// GenerateServiceToken for client.
	GenerateServiceToken string

	// VerifyServiceToken for client.
	VerifyServiceToken string
)

// RunCommandParams for client.
type RunCommandParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *zap.Logger
	Token     *Token
}

// RunCommand for client.
func RunCommand(params RunCommandParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return multierr.Append(generate(ctx, params.Token, params.Logger), verify(ctx, params.Token, params.Logger))
		},
	})
}

func generate(ctx context.Context, token *Token, logger *zap.Logger) error {
	p := strings.Split(GenerateServiceToken, ":")
	if len(p) != 2 {
		return nil
	}

	g := token.Generator(p[0], p[1])

	_, t, err := g.Generate(ctx)
	if err != nil {
		return err
	}

	logger.Info("generated service token", zap.String("token", string(t)))

	return nil
}

func verify(ctx context.Context, token *Token, logger *zap.Logger) error {
	p := strings.Split(VerifyServiceToken, ":")
	if len(p) != 4 {
		return nil
	}

	v := token.Verifier(p[0], p[1], p[2])

	if _, err := v.Verify(ctx, []byte(p[3])); err != nil {
		return err
	}

	logger.Info("verified service token")

	return nil
}
