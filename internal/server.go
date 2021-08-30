package server

import (
	"go.uber.org/fx"
	"net/http"
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
