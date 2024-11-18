package main

import (
	"context"
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
	envTests       = "tests"
)

const (
	defaultLogPath = "/root/logs/auth-service.log"
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
	case envTests:
		logFile, err := os.OpenFile(defaultLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic("Could not open log file: " + err.Error())
		}

		log = slog.New(NewMultiHandler(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
		))
	}

	return log
}

type MultiHandler struct {
	handlers []slog.Handler
}

func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

func (m *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, handler := range m.handlers {
		if err := handler.Handle(ctx, r); err != nil {
			return err
		}
	}
	return nil
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range m.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var newHandlers []slog.Handler
	for _, handler := range m.handlers {
		newHandlers = append(newHandlers, handler.WithAttrs(attrs))
	}
	return &MultiHandler{handlers: newHandlers}
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	var newHandlers []slog.Handler
	for _, handler := range m.handlers {
		newHandlers = append(newHandlers, handler.WithGroup(name))
	}
	return &MultiHandler{handlers: newHandlers}
}
