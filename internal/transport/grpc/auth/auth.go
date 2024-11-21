package authgrpc

import (
	"context"
	authpb "github.com/repooooo/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Auth defines the authentication interface.
type Auth interface {
	Login(ctx context.Context, username, password string) (bool, string, string, error)
	Logout(ctx context.Context, token string) (bool, string, error)
}

// serverAPI implements the AuthServiceServer interface.
type serverAPI struct {
	authpb.UnimplementedAuthServiceServer
	auth Auth
}

// newServerAPI creates a new instance of serverAPI.
func newServerAPI(auth Auth) *serverAPI {
	return &serverAPI{
		auth: auth,
	}
}

// Register registers the AuthServiceServer with gRPC.
func Register(gRPC *grpc.Server, auth Auth) {
	authpb.RegisterAuthServiceServer(gRPC, newServerAPI(auth))
	RegisterHealthCheck(gRPC)

	// TODO: Not secure. != production
	reflection.Register(gRPC)
}

// RegisterHealthCheck register health-check service.
func RegisterHealthCheck(gRPC *grpc.Server) {
	const serviceName = "auth_service"

	healthServer := health.NewServer()

	healthServer.SetServingStatus(serviceName, healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(gRPC, healthServer)
}
