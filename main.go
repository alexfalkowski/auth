package main

import (
	"os"

	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/cmd"
	"github.com/alexfalkowski/auth/rotate"
	sc "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	command := sc.New(cmd.Version)
	command.AddServer(cmd.ServerOptions...)

	cl := command.AddClient(cmd.ClientOptions...)
	sc.StringVar(cl, client.GenerateServiceToken, "generate", "g", "", "generate a service token")
	sc.StringVar(cl, client.VerifyServiceToken, "verify", "v", "", "verify a service token")

	ro := command.AddClientCommand("rotate", "Regenerate an existing configuration.", cmd.RotateOptions...)
	sc.StringVar(ro, sc.OutputFlag, "output", "o", "env:ROTATE_CONFIG_FILE", "output config location (format kind:location, default env:ROTATE_CONFIG_FILE)")
	sc.BoolVar(ro, rotate.Admins, "admins", "a", false, "admins configuration")
	sc.BoolVar(ro, rotate.Services, "services", "s", false, "services configuration")

	return command
}
