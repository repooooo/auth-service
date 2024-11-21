package http

import (
	"context"
	"errors"
	"fmt"
	graphhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/repooooo/auth-service/internal/transport/http/graph/auth"
	"github.com/repooooo/auth-service/internal/transport/http/rest/handler"
	"github.com/repooooo/go-utils/sl"
	"github.com/repooooo/graphqls/go/gen/auth"
	authresolver "github.com/repooooo/graphqls/go/resolvers/auth"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type App struct {
	log        *slog.Logger
	httpServer *http.Server
	port       int
}

type Auth interface {
	Login(
		ctx context.Context,
		username string,
		password string,
	) (
		success bool,
		message string,
		token string,
		err error,
	)
	Logout(
		ctx context.Context,
		token string,
	) (
		success bool,
		message string,
		err error,
	)
}

func New(
	log *slog.Logger,
	authService Auth,
	port int,
) *App {
	// transport.http.rest.handler
	http.HandleFunc("/health", handler.HealthCheck)

	// transport.http.graph.auth
	graphAuth := graphhandler.NewDefaultServer(authgen.NewExecutableSchema(
		authgen.Config{Resolvers: authresolver.NewResolver(authgraph.NewAuthHandler(authService))}))

	http.Handle("/auth", graphAuth)

	// GraphQL playground
	http.Handle(
		"/playground",
		playground.Handler(
			"GraphQL playground for auth",
			"/auth",
		),
	)

	httpServer := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &App{
		log:        log,
		httpServer: httpServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const operation = "app.http.Run"
	log := a.log.With(
		slog.String("operation", operation),
		slog.Int("port", a.port),
	)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error starting http server", sl.Err(err))
		}
	}()

	select {
	case <-time.After(1 * time.Second):
		log.Info("http server is running", slog.String("address", a.httpServer.Addr))
	case <-time.After(10 * time.Second):
		log.Error("http server failed to start in time")
		return fmt.Errorf("%s: server did not start in time", operation)
	}

	wg.Wait()

	return nil
}

func (a *App) Stop() {
	const operation = "app.http.Stop"

	a.log.With(slog.String("operation", operation)).
		Info("stopping http server")

	// TODO: handle error
	_ = a.httpServer.Shutdown(context.Background())
}
