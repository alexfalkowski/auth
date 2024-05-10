package main

import (
	"os"

	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/cmd"
	"github.com/alexfalkowski/auth/rotate"
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
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
	flags.StringVar(cl, client.GenerateFlag, "generate", "g", "", "generate a service token")
	flags.StringVar(cl, client.VerifyFlag, "verify", "v", "", "verify a service token")

	ro := command.AddClientCommand("rotate", "Regenerate an existing configuration.", cmd.RotateOptions...)
	flags.StringVar(ro, sc.OutputFlag, "output", "o", "env:ROTATE_CONFIG_FILE", "output config location (format kind:location, default env:ROTATE_CONFIG_FILE)")
	flags.BoolVar(ro, rotate.AdminsFlag, "admins", "a", false, "admins configuration")
	flags.BoolVar(ro, rotate.ServicesFlag, "services", "s", false, "services configuration")

	return command
}
