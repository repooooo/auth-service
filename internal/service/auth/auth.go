package auth

import (
	"context"
	"fmt"
	"github.com/repooooo/go-utils/sl"
	"log/slog"
)

type Auth struct {
	log      *slog.Logger
	provider Provider
}

type Provider interface {
	Login(ctx context.Context, username string, password string) (success bool, message string, token string, err error)
	Logout(ctx context.Context, token string) (success bool, message string, err error)
}

func New(log *slog.Logger, provider Provider) *Auth {
	return &Auth{
		log:      log,
		provider: provider,
	}
}

func (a *Auth) Login(ctx context.Context, username string, password string) (bool, string, string, error) {
	const operation = "Auth.Login"

	log := a.log.With(
		slog.String("operation", operation),
		slog.String("username", username),
	)

	log.Info("attempting to login")

	success, message, token, err := a.provider.Login(ctx, username, password)
	if err != nil {
		log.Error("failed to login", sl.Err(err))

		return false, "", "", fmt.Errorf("%s: %w", operation, err)
	}

	log.Info("login success", slog.String("message", fmt.Sprint(message)))

	return success, message, token, nil
}

func (a *Auth) Logout(ctx context.Context, token string) (bool, string, error) {
	const operation = "Auth.Logout"

	log := a.log.With(
		slog.String("operation", operation),
	)

	log.Info("attempting to logout")

	success, message, err := a.provider.Logout(ctx, token)
	if err != nil {
		log.Error("failed to logout", sl.Err(err))

		return false, "", fmt.Errorf("%s: %w", operation, err)
	}

	log.Info("logout success", slog.String("message", message))

	return success, message, nil
}
