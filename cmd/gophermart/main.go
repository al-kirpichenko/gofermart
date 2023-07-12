package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	"github.com/al-kirpichenko/gofermart/internal/api"
	"github.com/al-kirpichenko/gofermart/internal/router"

	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {

	cfg := config.NewConfig()
	server := api.NewServer(cfg)

	r := router.Router(server)

	srv := &http.Server{
		Addr:    cfg.ServiceHost,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //nolint:govet
	<-quit
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	if <-ctx.Done(); true {
		zap.L().Info("Shutdown: timeout of 5 seconds.")
	}
}
