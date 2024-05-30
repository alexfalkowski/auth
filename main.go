package main

import (
	"os"

	"github.com/alexfalkowski/auth/cmd"
	"github.com/alexfalkowski/auth/cmd/rotate"
	"github.com/alexfalkowski/auth/cmd/token"
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:AUTH_CONFIG_FILE")
	c.AddServer(cmd.ServerOptions...)

	cl := c.AddClientCommand("token", "Perform actions with tokens.", cmd.TokenOptions...)
	flags.StringVar(cl, token.GenerateFlag, "generate", "g", "", "generate a service token")
	flags.StringVar(cl, token.VerifyFlag, "verify", "v", "", "verify a service token")

	ro := c.AddClientCommand("rotate", "Regenerate an existing configuration.", cmd.RotateOptions...)
	c.RegisterOutput(ro, "env:AUTH_APP_CONFIG_FILE")
	flags.BoolVar(ro, rotate.AdminsFlag, "admins", "a", false, "admins configuration")
	flags.BoolVar(ro, rotate.ServicesFlag, "services", "s", false, "services configuration")

	return c
}
