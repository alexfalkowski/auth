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
			ks := func(ctx context.Context) (context.Context, error) {
				r, err := generateKeys(ctx, params)

				return context.WithValue(ctx, k("r"), r), err
			}

			as := func(ctx context.Context) (context.Context, error) {
				return ctx, generateAdmins(ctx, params)
			}

			ss := func(ctx context.Context) (context.Context, error) {
				r := ctx.Value(k("r")).(*key.RSA)

				return ctx, generateServices(ctx, r, params)
			}

			c := func(ctx context.Context) (context.Context, error) {
				m, err := params.Factory.Create(params.OutputConfig.Kind())

				return context.WithValue(ctx, k("m"), m), err
			}

			m := func(ctx context.Context) (context.Context, error) {
				m := ctx.Value(k("m")).(marshaller.Marshaller)
				d, err := m.Marshal(params.Config)

				return context.WithValue(ctx, k("d"), d), err
			}

			o := func(ctx context.Context) (context.Context, error) {
				d := ctx.Value(k("d")).([]byte)

				return ctx, params.OutputConfig.Write(d, fs.FileMode(0o600))
			}

			_, err := fold(ctx, ks, as, ss, c, m, o)

			return err
		},
	})
}

func generateKeys(ctx context.Context, params RunCommandParams) (*key.RSA, error) {
	if !isAll() {
		return rsa(params.Config.Key.RSA.Public, params.Config.Key.RSA.Private)
	}

	rs := func(ctx context.Context) (context.Context, error) {
		pub, pri, err := params.KeyGenerator.Generate("rsa")
		ctx = context.WithValue(ctx, k("pub"), pub)
		ctx = context.WithValue(ctx, k("pri"), pri)

		return ctx, err
	}

	rsa := func(ctx context.Context) (context.Context, error) {
		r, err := rsa(ctx.Value(k("pub")).(string), ctx.Value(k("pri")).(string))

		return context.WithValue(ctx, k("r"), r), err
	}

	ars := func(ctx context.Context) (context.Context, error) {
		params.Config.Key.RSA.Public = ctx.Value(k("pub")).(string)
		params.Config.Key.RSA.Private = ctx.Value(k("pri")).(string)

		return ctx, nil
	}

	es := func(ctx context.Context) (context.Context, error) {
		pub, pri, err := params.KeyGenerator.Generate("ed25519")
		ctx = context.WithValue(ctx, k("pub"), pub)
		ctx = context.WithValue(ctx, k("pri"), pri)

		return ctx, err
	}

	aes := func(ctx context.Context) (context.Context, error) {
		params.Config.Key.Ed25519.Public = ctx.Value(k("pub")).(string)
		params.Config.Key.Ed25519.Private = ctx.Value(k("pri")).(string)

		return ctx, nil
	}

	ctx, err := fold(ctx, rs, rsa, ars, es, aes)
	if err != nil {
		return nil, err
	}

	return ctx.Value(k("r")).(*key.RSA), nil
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
	if !Admins && !isAll() {
		return nil
	}

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
	if !Services && !isAll() {
		return nil
	}

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
