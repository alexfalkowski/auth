package rotate

import (
	"context"
	"io/fs"

	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/auth/key"
	"github.com/alexfalkowski/auth/password"
	"github.com/alexfalkowski/go-service/cmd"
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

// Params for rotate.
type Params struct {
	fx.In

	Lifecycle    fx.Lifecycle
	OutputConfig *cmd.OutputConfig
	Key          *key.Generator
	Secure       *password.Secure
	Map          *marshaller.Map
	Config       *config.Config
	RSA          rsa.Algo
	Logger       *zap.Logger
}

// Start for rotate.
func Start(params Params) {
	cmd.Start(params.Lifecycle, func(_ context.Context) {
		generateAdmins(params)
		generateServices(params)

		m := params.Map.Get(params.OutputConfig.Kind())

		d, err := m.Marshal(params.Config)
		runtime.Must(err)

		runtime.Must(params.OutputConfig.Write(d, fs.FileMode(0o600)))
	})
}

func generateAdmins(params Params) {
	if !flags.IsSet(AdminsFlag) {
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

func generateServices(params Params) {
	if !flags.IsSet(ServicesFlag) {
		return
	}

	for _, svc := range params.Config.Server.V1.Services {
		p, err := params.Secure.Generate(password.DefaultLength)
		runtime.Must(err)

		h, err := params.Secure.Hash(p)
		runtime.Must(err)

		b, err := params.RSA.Encrypt(p)
		runtime.Must(err)

		svc.Hash = h

		params.Logger.Info("generated service password/token", zap.String("id", svc.ID), zap.String("password", p), zap.String("token", b))
	}
}
