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
	cl.PersistentFlags().StringVarP(&client.GenerateServiceToken, "generate", "g", "", "generate a service token")
	cl.PersistentFlags().StringVarP(&client.VerifyServiceToken, "verify", "v", "", "verify a service token")

	ro := command.AddClientCommand("rotate", "Regenerate an existing configuration.", cmd.RotateOptions...)
	ro.PersistentFlags().StringVarP(&rotate.OutputFlag, "output", "o", "env:ROTATE_CONFIG_FILE", "output config location (format kind:location, default env:ROTATE_CONFIG_FILE)")
	ro.PersistentFlags().BoolVarP(&rotate.Admins, "admins", "a", false, "admins configuration")
	ro.PersistentFlags().BoolVarP(&rotate.Services, "services", "s", false, "services configuration")

	return command
}
