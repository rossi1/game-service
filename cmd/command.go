package cmd

import (
	"game-service/config"
	"game-service/server"

	"github.com/spf13/cobra"
)

func runServerCMD(dep config.Deps) (*cobra.Command, error) {
	var serverCmd = &cobra.Command{
		Use:   "runserver",
		Short: "run the web server",
	}

	serverCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := server.RunServer(cmd.Context(), dep); err != nil {
			return err
		}
		return nil

	}

	return serverCmd, nil
}
