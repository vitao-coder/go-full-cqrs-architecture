package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/vitao-coder/go-full-cqrs-architecture/configuration"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/logging"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

const (
	timeout = 60
)

type ServerInput struct {
	fx.In
	Endpoints []HTTPEndpoint `group:"endpoints"`
}

type HTTPEndpoint interface {
	http.Handler
	HttpMethod() string
	HttpPath() string
}

func NewConfiguration() configuration.Configuration {
	f, err := os.Open("configuration/config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg configuration.Configuration
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func NewServer(logger logging.Logger, input ServerInput) *chi.Mux {
	logger.Info("Starting registering server...")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(timeout * time.Second))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	logger.Info("Registering endpoints...")
	r.Route("/go-full-app", func(r chi.Router) {
		for _, endpoint := range input.Endpoints {
			r.Method(endpoint.HttpMethod(), endpoint.HttpPath(), endpoint)
		}
	})
	logger.Info("Endpoints registered...")
	return r
}

func StartServer(lc fx.Lifecycle, logger logging.Logger, server *chi.Mux, config configuration.Configuration) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting server")
			go http.ListenAndServe(":"+config.Server.Port, server)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping server")
			return nil
		},
	})
}

func Serve() {
	ServerModule := fx.Provide(
		NewConfiguration,
		NewServer,
	)

	app := fx.New(fx.Options(
		ServerModule,
		PackagesModule,
	),
		fx.Invoke(StartServer))

	app.Run()
}
