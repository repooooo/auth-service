package authgraph

import "context"

type Auth interface {
	Login(ctx context.Context, username, password string) (bool, string, string, error)
	Logout(ctx context.Context, token string) (bool, string, error)
}

type Handler struct {
	auth Auth
}

func NewHandler(auth Auth) *Handler {
	return &Handler{auth: auth}
}
