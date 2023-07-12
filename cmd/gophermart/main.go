package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	"github.com/al-kirpichenko/gofermart/internal/api"
	"github.com/al-kirpichenko/gofermart/internal/router"
)

func main() {

	cfg := config.NewConfig()
	server := api.NewServer(cfg)

	r := router.Router(server)

	srv := &http.Server{
		Addr:    cfg.ServiceHost,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
}
