package main

import (
	"github.com/repooooo/auth-service/internal/app"
	"github.com/repooooo/auth-service/internal/config"
	"github.com/repooooo/go-utils/loader"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal       = "local"
	envDevelopment = "development"
	envProduction  = "production"
)

func main() {
	// Load configuration
	cfg := config.New()
	configPath := loader.FetchConfigPath()
	cl := loader.NewConfigLoader(configPath)
	cl.MustLoad(cfg)

	// Set up logger
	log := setupLogger(cfg.Env)

	// Initialize application with configuration
	application := app.New(
		log,
		cfg.GRPC.Port,
		cfg.DSN,
		cfg.HTTP.Port,
	)

	// Start gRPC and HTTP servers in separate goroutines
	go application.GRPCServer.MustRun()
	go application.HTTPServer.MustRun()

	// Listen for termination signals (e.g., SIGTERM, SIGINT)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Wait for signal, then stop servers gracefully
	signalStop := <-stop
	log.Info("stopping application", slog.String("signal", signalStop.String()))
	application.GRPCServer.Stop()
	application.HTTPServer.Stop()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case envDevelopment:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case envProduction:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}

	return log
}
