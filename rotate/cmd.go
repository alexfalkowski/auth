package rotate

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"io/fs"

	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/marshaller"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// OutputFlag for rotate.
var OutputFlag string

// OutputConfig for rotate.
type OutputConfig struct {
	*cmd.Config
}

// NewOutputConfig for rotate.
func NewOutputConfig(factory *marshaller.Factory) (*OutputConfig, error) {
	c, err := cmd.NewConfig(OutputFlag, factory)
	if err != nil {
		return nil, err
	}

	return &OutputConfig{Config: c}, nil
}

// RunCommandParams for client.
type RunCommandParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	OutputConfig *OutputConfig
	KeyGenerator *key.Generator
	Secure       *password.Secure
	Factory      *marshaller.Factory
	Config       *config.Config
	Logger       *zap.Logger
}

// RunCommand for client.
func RunCommand(params RunCommandParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			public, private, err := params.KeyGenerator.Generate("rsa")
			if err != nil {
				return err
			}

			params.Config.Key.RSA.Public = public
			params.Config.Key.RSA.Private = private

			r, err := rsa(public, private)
			if err != nil {
				return err
			}

			public, private, err = params.KeyGenerator.Generate("ed25519")
			if err != nil {
				return err
			}

			params.Config.Key.Ed25519.Public = public
			params.Config.Key.Ed25519.Private = private

			if err := generateAdmins(ctx, params); err != nil {
				return err
			}

			if err := generateServices(ctx, r, params); err != nil {
				return err
			}

			m, err := params.Factory.Create(params.OutputConfig.Kind())
			if err != nil {
				return err
			}

			d, err := m.Marshal(params.Config)
			if err != nil {
				return err
			}

			return params.OutputConfig.Write(d, fs.FileMode(0o600))
		},
	})
}

func rsa(public, private string) (*key.RSA, error) {
	k, err := base64.StdEncoding.DecodeString(private)
	if err != nil {
		return nil, err
	}

	pk, err := x509.ParsePKCS1PrivateKey(k)
	if err != nil {
		return nil, err
	}

	k, err = base64.StdEncoding.DecodeString(public)
	if err != nil {
		return nil, err
	}

	puk, err := x509.ParsePKCS1PublicKey(k)
	if err != nil {
		return nil, err
	}

	return key.NewRSA(puk, pk), nil
}

func generateAdmins(ctx context.Context, params RunCommandParams) error {
	for _, admin := range params.Config.Server.V1.Admins {
		p, err := params.Secure.Generate(ctx, password.DefaultLength)
		if err != nil {
			return err
		}

		h, err := params.Secure.Hash(ctx, p)
		if err != nil {
			return err
		}

		admin.Hash = h

		params.Logger.Info("generated admin password", zap.String("id", admin.ID), zap.String("password", p))
	}

	return nil
}

func generateServices(ctx context.Context, rsa *key.RSA, params RunCommandParams) error {
	for _, svc := range params.Config.Server.V1.Services {
		p, err := params.Secure.Generate(ctx, password.DefaultLength)
		if err != nil {
			return err
		}

		h, err := params.Secure.Hash(ctx, p)
		if err != nil {
			return err
		}

		b, err := rsa.Encrypt(ctx, p)
		if err != nil {
			return err
		}

		svc.Hash = h

		params.Logger.Info("generated service password/token", zap.String("id", svc.ID), zap.String("password", p), zap.String("token", b))
	}

	return nil
}
