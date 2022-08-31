package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"game-service/cmd"
	"game-service/config"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.GetConfig(".")
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to boot up configuration: %s", err.Error()))
		os.Exit(1)
	}

	deps, err := config.GetDeps(ctx, cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to get dependencies: %s", err.Error()))
		os.Exit(1)
	}

	err = cmd.Execute(ctx, deps)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

}
