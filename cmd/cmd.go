package cmd

import (
	"context"
	"fmt"
	"game-service/config"

	"github.com/spf13/cobra"
)

func Execute(ctx context.Context, dep config.Deps) error {
	var CMD = &cobra.Command{
		Use:     "Start up server",
		Short:   "Start application server",
		Long:    "game service application",
		Version: "v0.1.0",
	}

	serverCMD, err := runServerCMD(dep)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	CMD.AddCommand(serverCMD)

	if err := CMD.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}
