package main

import (
	"os"

	"github.com/alexfalkowski/auth/cmd"
	scmd "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	command := scmd.New()

	command.AddServer(cmd.ServerOptions)
	command.AddVersion(cmd.Version)

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
