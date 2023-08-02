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
	"github.com/al-kirpichenko/gofermart/internal/services/scheduler"

	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()

	cfg := config.NewConfig()
	server := api.NewServer(cfg)
	server.Logger = logger

	srv := &http.Server{
		Addr:    cfg.ServiceHost,
		Handler: router.Router(server),
	}

	// Graceful shutdown

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// запускаем опрос  заказов со статусом PROCESSING
	go scheduler.UpdateOrders(server)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Graceful shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Graceful shutdown server: ", zap.Error(err))
	}

	if <-ctx.Done(); true {
		logger.Info("Graceful shutdown: timeout of 5 seconds.")
	}
}
