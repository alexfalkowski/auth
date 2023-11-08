package client

import (
	"context"
	"strings"

	"go.uber.org/fx"
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
			if ok, kind, audience := generateServiceToken(); ok {
				g := params.Token.Generator(kind, audience)

				_, t, err := g.Generate(ctx)
				if err != nil {
					return err
				}

				params.Logger.Info("generated service token", zap.String("token", string(t)))
			}

			if ok, kind, audience, action, token := verifyServiceToken(); ok {
				v := params.Token.Verifier(kind, audience, action)

				if _, err := v.Verify(ctx, []byte(token)); err != nil {
					return err
				}

				params.Logger.Info("verified service token")
			}

			return nil
		},
	})
}

func generateServiceToken() (bool, string, string) {
	p := strings.Split(GenerateServiceToken, ":")
	if len(p) != 2 {
		return false, "", ""
	}

	return true, p[0], p[1]
}

func verifyServiceToken() (bool, string, string, string, string) {
	p := strings.Split(VerifyServiceToken, ":")
	if len(p) != 4 {
		return false, "", "", "", ""
	}

	return true, p[0], p[1], p[2], p[3]
}
