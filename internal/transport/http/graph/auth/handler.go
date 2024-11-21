package authgraph

import (
	"context"
	"github.com/repooooo/graphqls/go/gen/auth/model"
	"google.golang.org/grpc/status"
	"net/http"
)

const (
	ErrUsernameRequired = "username is required"
	ErrPasswordRequired = "password is required"
	ErrTokenRequired    = "token is required"
	ErrInternalServer   = "internal server error"
)

func (h *Handler) Login(ctx context.Context, input authmodel.LoginRequest) (*authmodel.LoginResponse, error) {
	if err := validateLoginRequest(input); err != nil {
		return nil, err
	}

	success, message, token, err := h.auth.Login(ctx, input.Username, input.Password)
	if err != nil {
		return nil, status.Error(http.StatusInternalServerError, ErrInternalServer)
	}

	return &authmodel.LoginResponse{
		Success: success,
		Message: message,
		Token:   token,
	}, nil
}

func validateLoginRequest(request authmodel.LoginRequest) error {
	if request.Username == "" {
		return status.Error(http.StatusBadRequest, ErrUsernameRequired)
	}

	if request.Password == "" {
		return status.Error(http.StatusBadRequest, ErrPasswordRequired)
	}

	return nil
}

func (h *Handler) Logout(ctx context.Context, input authmodel.LogoutRequest) (*authmodel.LogoutResponse, error) {
	if err := validateLogoutRequest(input); err != nil {
		return nil, err
	}

	// TODO h.auth.Logout

	return &authmodel.LogoutResponse{
		Success: true,
		Message: "Successfully logged out",
	}, nil
}

func validateLogoutRequest(request authmodel.LogoutRequest) error {
	if request.Token == "" {
		return status.Error(http.StatusBadRequest, ErrTokenRequired)
	}

	return nil
}
