package main

import (
	"os"

	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/cmd"
	"github.com/alexfalkowski/auth/rotate"
	scmd "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *scmd.Command {
	command := scmd.New(cmd.Version)

	command.AddServer(cmd.ServerOptions...)

	cl := command.AddClient(cmd.ClientOptions...)
	cl.PersistentFlags().StringVar(&client.GenerateServiceToken, "generate-service-token", "", "generate a service token")
	cl.PersistentFlags().StringVar(&client.VerifyServiceToken, "verify-service-token", "", "verify a service token")

	ro := command.AddClientCommand("rotate", "Regenerate an existing configuration.", cmd.RotateOptions...)
	ro.PersistentFlags().StringVar(
		&rotate.OutputFlag,
		"output", "env:ROTATE_CONFIG_FILE", "output config location (format kind:location, default env:ROTATE_CONFIG_FILE)",
	)
	ro.PersistentFlags().BoolVar(&rotate.Admins, "admins", false, "admins configuration")
	ro.PersistentFlags().BoolVar(&rotate.Services, "services", false, "services configuration")

	return command
}
