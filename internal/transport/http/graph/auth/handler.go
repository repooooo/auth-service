package authgraph

import (
	"context"
	"github.com/repooooo/graphqls/go/gen/auth/model"
)

type Handler interface {
	Login(ctx context.Context, input authmodel.LoginRequest) (*authmodel.LoginResponse, error)
	Logout(ctx context.Context, input authmodel.LogoutRequest) (*authmodel.LogoutResponse, error)
}
