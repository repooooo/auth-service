package main

import (
	"github.com/repooooo/auth/internal/app"
	"github.com/repooooo/auth/internal/config"
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
	cfg := config.New()

	configPath := loader.FetchConfigPath()
	cl := loader.NewConfigLoader(configPath)
	cl.MustLoad(cfg)

	log := setupLogger(cfg.Env)

	application := app.New(log, cfg.GRPC.Port, cfg.DSN)
	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signalStop := <-stop
	log.Info("stopping application", slog.String("signal", signalStop.String()))
	application.GRPCServer.Stop()
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
