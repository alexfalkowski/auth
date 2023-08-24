package main

import (
	"os"

	"github.com/alexfalkowski/auth/client"
	"github.com/alexfalkowski/auth/cmd"
	scmd "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *scmd.Command {
	command := scmd.New()

	command.AddServer(cmd.ServerOptions)

	c := command.AddClient(cmd.ClientOptions)
	c.PersistentFlags().Int32Var(
		&client.GenerateAccessToken,
		"generate-access-token", -1, "generate an access token",
	)
	c.PersistentFlags().StringVar(
		&client.GenerateServiceToken,
		"generate-service-token", "", "generate a service token",
	)
	c.PersistentFlags().StringVar(
		&client.VerifyServiceToken,
		"verify-service-token", "", "verify a service token",
	)

	command.AddVersion(cmd.Version)

	return command
}
