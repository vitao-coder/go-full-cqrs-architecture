package zap

import (
	"github.com/vitao-coder/go-full-cqrs-architecture/constants"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/logging"

	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(enviroment string) (logging.Logger, error) {
	var logger *zap.Logger
	var err error
	if enviroment == constants.ProductionEnviroment {
		logger, err = zap.NewProduction()
	} else if enviroment == constants.DevelopmentEnviroment {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return &zapLogger{logger: logger}, nil
}

func (zl zapLogger) Debug(msg string) {
	zl.logger.Debug(msg)
}

func (zl zapLogger) Info(msg string) {
	zl.logger.Info(msg)
}

func (zl zapLogger) Warn(msg string) {
	zl.logger.Warn(msg)
}

func (zl zapLogger) Error(msg string, err error) {
	zl.logger.Error(msg, zap.Error(err))
}
