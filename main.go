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
	c := sc.New(cmd.Version)
	c.RegisterInput("env:AUTH_IN_CONFIG_FILE")
	c.RegisterOutput("env:AUTH_OUT_CONFIG_FILE")
	c.AddServer(cmd.ServerOptions...)

	cl := c.AddClient(cmd.ClientOptions...)
	flags.StringVar(cl, client.GenerateFlag, "generate", "g", "", "generate a service token")
	flags.StringVar(cl, client.VerifyFlag, "verify", "v", "", "verify a service token")

	ro := c.AddClientCommand("rotate", "Regenerate an existing configuration.", cmd.RotateOptions...)
	flags.BoolVar(ro, rotate.AdminsFlag, "admins", "a", false, "admins configuration")
	flags.BoolVar(ro, rotate.ServicesFlag, "services", "s", false, "services configuration")

	return c
}
