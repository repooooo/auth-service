package suite

import (
	"context"
	"github.com/repooooo/auth-service/internal/config"
	"github.com/repooooo/go-utils/loader"
	authpb "github.com/repooooo/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strconv"
	"testing"
)

type Suite struct {
	*testing.T
	Cfg               *config.Config
	AuthServiceClient authpb.AuthServiceClient
}

const (
	grpcHost          = "localhost"
	configPathDefault = "./test-config/config.yaml"
)

// New creates new test suite.
//
// TODO: for pipeline tests we need to wait for app is ready
func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.New()

	cl := loader.NewConfigLoader(configPathDefault)
	cl.MustLoad(cfg)

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(context.Background(),
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:                 t,
		Cfg:               cfg,
		AuthServiceClient: authpb.NewAuthServiceClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
