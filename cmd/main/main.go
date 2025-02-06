package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/UsedC/packing-service/internal/app"
	"github.com/UsedC/packing-service/internal/calculator"
	"github.com/UsedC/packing-service/internal/config"
	"github.com/UsedC/packing-service/internal/server"
	"github.com/UsedC/packing-service/internal/service"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: config.LogLevel}))
	slog.SetDefault(logger)

	packService := service.NewPackService(config.PackSizes, calculator.CalculatePacks)

	app := app.NewApp(packService)

	handler := server.NewHandler(app)

	mux := server.NewRouter(handler)

	srv := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("server error", slog.String("error", err.Error()))
			quit <- syscall.SIGTERM
		}
	}()

	logger.Info("server started", slog.String("address", config.Address))

	<-quit

	logger.Info("server stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown error", slog.String("error", err.Error()))
	}

	logger.Info("server stopped")
}
