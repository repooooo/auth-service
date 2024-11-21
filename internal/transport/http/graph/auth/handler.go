package authgraph

import (
	"context"
	"github.com/repooooo/graphqls/go/gen/auth/model"
)

type Handler struct {
	auth Auth
}

func NewHandler(auth Auth) *Handler {
	return &Handler{auth: auth}
}

func (h *Handler) Login(ctx context.Context, input authmodel.LoginRequest) (*authmodel.LoginResponse, error) {
	//TODO: validator for request
	success, message, token, err := h.auth.Login(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	return &authmodel.LoginResponse{
		Success: success,
		Message: message,
		Token:   token,
	}, nil
}

func (h *Handler) Logout(ctx context.Context, input authmodel.LogoutRequest) (*authmodel.LogoutResponse, error) {
	return &authmodel.LogoutResponse{
		Success: true,
		Message: "Successfully logged out",
	}, nil
}
