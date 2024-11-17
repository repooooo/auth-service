package app

import (
	grpcapp "github.com/repooooo/auth-service/internal/app/grpc"
	"github.com/repooooo/auth-service/internal/service/auth"
	"github.com/repooooo/auth-service/internal/storage/postgres"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	dsn string,
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

	return &App{
		GRPCServer: grpcApp,
	}
}
