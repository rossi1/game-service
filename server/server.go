package server

import (
	"context"
	"fmt"
	"game-service/config"
	"game-service/server/handler"
	"log"
	"net/http"
	"time"
)

func RunServer(ctx context.Context, deps config.Deps) error {
	port := deps.Port
	host := deps.Host
	address := fmt.Sprintf("%s:%s", host, port)

	hnd := handler.NewHandler()

	httpServer := &http.Server{
		Handler:           hnd.GetRouter(),
		Addr:              address,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()
	log.Println("Shutting down server....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {

		cancel()

	}()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return nil

}
