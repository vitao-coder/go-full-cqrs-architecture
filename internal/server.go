package server

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
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

func NewLogger() *zap.Logger{
	logger, err := zap.NewProduction()

	if err!=nil{
		panic(err)
	}

	return logger
}
