package postgres

import (
	"context"
	"log/slog"
)

func (s *Storage) Login(ctx context.Context, username string, password string) (bool, string, string, error) {
	const operation = "storage.postgres.Login"

	log := s.log.With(
		slog.String("operation", operation),
		slog.String("username", username),
	)
	// TODO:
	log.Info("todo login")

	return false, "", "", nil
}

func (s *Storage) Logout(ctx context.Context, token string) (success bool, message string, err error) {
	const operation = "storage.postgres.Logout"

	return false, "", nil
}

func (s *Storage) Create(ctx context.Context, username, password string) (success bool, message string, err error) {
	const operation = "storage.postgres.Create"

	return false, "", nil
}
