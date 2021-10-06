package server

import (
	"github.com/vitao-coder/go-full-cqrs-architecture/configuration"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/database"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/database/mongodb"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/logging"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/logging/zap"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/messaging"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/messaging/pulsar"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/metrics"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/metrics/prometheus"

	"go.uber.org/fx"
)

var PackagesModule = fx.Provide(
	NewDatabase,
	NewMessaging,
	NewLogger,
	NewMetrics,
)

func NewDatabase(config configuration.Configuration) (database.Database, error) {
	return mongodb.NewMongoDatabase(config.Database.ConnectionString, config.Database.Database)
}

func NewMessaging(config configuration.Configuration) (messaging.Messaging, error) {
	return pulsar.NewPulsarClient(config.Messaging.URL)
}

func NewLogger(config configuration.Configuration) (logging.Logger, error) {
	return zap.NewZapLogger(config.Server.Enviroment)
}

func NewMetrics(logger logging.Logger) (metrics.Metrics, error) {
	return prometheus.NewPrometheusService(logger)
}
