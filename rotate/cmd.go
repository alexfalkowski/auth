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
	"github.com/alexfalkowski/go-service/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// OutputFlag for rotate.
	OutputFlag string

	// Admins to be rotated.
	Admins bool

	// Services to be rotated.
	Services bool
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
	Key          *key.Generator
	Secure       *password.Secure
	Factory      *marshaller.Factory
	Config       *config.Config
	Logger       *zap.Logger
}

// RunCommand for client.
func RunCommand(params RunCommandParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = runtime.Convert(r)
				}
			}()

			r := generateKeys(params)

			generateAdmins(ctx, params)
			generateServices(ctx, r, params)

			m, err := params.Factory.Create(params.OutputConfig.Kind())
			runtime.Must(err)

			d, err := m.Marshal(params.Config)
			runtime.Must(err)

			runtime.Must(params.OutputConfig.Write(d, fs.FileMode(0o600)))

			return
		},
	})
}

func isAll() bool {
	return !Admins && !Services
}

func generateKeys(params RunCommandParams) *key.RSA {
	if !isAll() {
		return rsa(params.Config.Key.RSA.Public, params.Config.Key.RSA.Private)
	}

	public, private, err := params.Key.Generate("rsa")
	runtime.Must(err)

	params.Config.Key.RSA.Public = public
	params.Config.Key.RSA.Private = private

	r := rsa(public, private)

	public, private, err = params.Key.Generate("ed25519")
	runtime.Must(err)

	params.Config.Key.Ed25519.Public = public
	params.Config.Key.Ed25519.Private = private

	return r
}

func rsa(public, private string) *key.RSA {
	k, err := base64.StdEncoding.DecodeString(private)
	runtime.Must(err)

	pk, err := x509.ParsePKCS1PrivateKey(k)
	runtime.Must(err)

	k, err = base64.StdEncoding.DecodeString(public)
	runtime.Must(err)

	puk, err := x509.ParsePKCS1PublicKey(k)
	runtime.Must(err)

	return key.NewRSA(puk, pk)
}

func generateAdmins(ctx context.Context, params RunCommandParams) {
	if !Admins && !isAll() {
		return
	}

	for _, admin := range params.Config.Server.V1.Admins {
		p, err := params.Secure.Generate(ctx, password.DefaultLength)
		runtime.Must(err)

		h, err := params.Secure.Hash(ctx, p)
		runtime.Must(err)

		admin.Hash = h

		params.Logger.Info("generated admin password", zap.String("id", admin.ID), zap.String("password", p))
	}
}

func generateServices(ctx context.Context, rsa *key.RSA, params RunCommandParams) {
	if !Services && !isAll() {
		return
	}

	for _, svc := range params.Config.Server.V1.Services {
		p, err := params.Secure.Generate(ctx, password.DefaultLength)
		runtime.Must(err)

		h, err := params.Secure.Hash(ctx, p)
		runtime.Must(err)

		b, err := rsa.Encrypt(ctx, p)
		runtime.Must(err)

		svc.Hash = h

		params.Logger.Info("generated service password/token", zap.String("id", svc.ID), zap.String("password", p), zap.String("token", b))
	}
}
