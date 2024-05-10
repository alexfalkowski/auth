package rotate

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/crypto/ed25519"
	"github.com/alexfalkowski/go-service/crypto/rsa"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// AdminsFlag to be rotated.
	AdminsFlag = flags.Bool()

	// ServicesFlag to be rotated.
	ServicesFlag = flags.Bool()
)

// RunCommandParams for client.
type RunCommandParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	OutputConfig *cmd.OutputConfig
	Key          *key.Generator
	Secure       *password.Secure
	Factory      *marshaller.Factory
	Config       *config.Config
	Ed25519      *ed25519.Config
	RSA          *rsa.Config
	Logger       *zap.Logger
}

// RunCommand for client.
func RunCommand(params RunCommandParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = runtime.ConvertRecover(r)
				}
			}()

			r := generateKeys(params)

			generateAdmins(params)
			generateServices(r, params)

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
	return !*AdminsFlag && !*ServicesFlag
}

func generateKeys(params RunCommandParams) *key.RSA {
	if !isAll() {
		return rsaKey(params.RSA.Public, params.RSA.Private)
	}

	public, private, err := params.Key.Generate("rsa")
	runtime.Must(err)

	params.RSA.Public = public
	params.RSA.Private = private

	r := rsaKey(public, private)

	public, private, err = params.Key.Generate("ed25519")
	runtime.Must(err)

	params.Ed25519.Public = public
	params.Ed25519.Private = private

	return r
}

func rsaKey(public, private string) *key.RSA {
	a, err := rsa.NewAlgo(&rsa.Config{Public: public, Private: private})
	runtime.Must(err)

	return key.NewRSA(a)
}

func generateAdmins(params RunCommandParams) {
	if !*AdminsFlag && !isAll() {
		return
	}

	for _, admin := range params.Config.Server.V1.Admins {
		p, err := params.Secure.Generate(password.DefaultLength)
		runtime.Must(err)

		h, err := params.Secure.Hash(p)
		runtime.Must(err)

		admin.Hash = h

		params.Logger.Info("generated admin password", zap.String("id", admin.ID), zap.String("password", p))
	}
}

func generateServices(rsa *key.RSA, params RunCommandParams) {
	if !*ServicesFlag && !isAll() {
		return
	}

	for _, svc := range params.Config.Server.V1.Services {
		p, err := params.Secure.Generate(password.DefaultLength)
		runtime.Must(err)

		h, err := params.Secure.Hash(p)
		runtime.Must(err)

		b, err := rsa.Encrypt(p)
		runtime.Must(err)

		svc.Hash = h

		params.Logger.Info("generated service password/token", zap.String("id", svc.ID), zap.String("password", p), zap.String("token", b))
	}
}
