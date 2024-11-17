package authgrpc

import "context"

// Auth defines the authentication interface.
type Auth interface {
	Login(ctx context.Context, username, password string) (bool, string, string, error)
	Logout(ctx context.Context, token string) (bool, string, error)
}
