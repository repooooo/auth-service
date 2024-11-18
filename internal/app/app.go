package app

import (
	grpcapp "github.com/repooooo/auth-service/internal/app/grpc"
	httpapp "github.com/repooooo/auth-service/internal/app/http"
	"github.com/repooooo/auth-service/internal/service/auth"
	"github.com/repooooo/auth-service/internal/storage/postgres"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
	HTTPServer *httpapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	dsn string,
	httpPort int,
) *App {
	storage, err := postgres.New(
		log,
		dsn,
	)
	if err != nil {
		panic(err)
	}

	authService := auth.New(
		log,
		storage,
	)

	grpcApp := grpcapp.New(
		log,
		authService,
		grpcPort,
	)

	httpApp := httpapp.New(
		log,
		httpPort,
	)

	return &App{
		GRPCServer: grpcApp,
		HTTPServer: httpApp,
	}
}
