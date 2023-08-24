package client

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// GenerateAccessToken for client.
	GenerateAccessToken int32

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
	Client    *Client
}

// RunCommand for client.
func RunCommand(params RunCommandParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if GenerateAccessToken >= 0 {
				token, err := params.Client.GenerateAccessToken(ctx, uint32(GenerateAccessToken))
				if err != nil {
					return err
				}

				params.Logger.Info("generated access token", zap.String("token", token))
			}

			if ok, kind, audience := generateServiceToken(); ok {
				token, err := params.Client.GenerateServiceToken(ctx, kind, audience)
				if err != nil {
					return err
				}

				params.Logger.Info("generated service token", zap.String("token", token))
			}

			if ok, kind, audience, action, token := verifyServiceToken(); ok {
				err := params.Client.VerifyServiceToken(ctx, token, kind, audience, action)
				if err != nil {
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
