package token

import (
	"context"
	"strings"

	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/runtime"
	"go.uber.org/fx"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

var (
	// GenerateFlag for token.
	GenerateFlag = flags.String()

	// VerifyFlag for token.
	VerifyFlag = flags.String()
)

// Start for token.
func Start(lc fx.Lifecycle, logger *zap.Logger, token *client.Token) {
	cmd.Start(lc, func(ctx context.Context) {
		err := multierr.Append(generate(ctx, token, logger), verify(ctx, token, logger))
		runtime.Must(err)
	})
}

func generate(ctx context.Context, token *client.Token, logger *zap.Logger) error {
	p := strings.Split(*GenerateFlag, ":")
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

func verify(ctx context.Context, token *client.Token, logger *zap.Logger) error {
	p := strings.Split(*VerifyFlag, ":")
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
