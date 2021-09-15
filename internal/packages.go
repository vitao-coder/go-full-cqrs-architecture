package server

import (
	"go-full-cqrs-architecture/configuration"
	"go-full-cqrs-architecture/packages/database"
	"go-full-cqrs-architecture/packages/database/mongodb"
	"go-full-cqrs-architecture/packages/logging"
	"go-full-cqrs-architecture/packages/logging/zap"
	"go-full-cqrs-architecture/packages/messaging"
	"go-full-cqrs-architecture/packages/messaging/pulsar"
	"go-full-cqrs-architecture/packages/metrics"
	"go-full-cqrs-architecture/packages/metrics/prometheus"

	"go.uber.org/fx"
)

var PackagesModule = fx.Provide(
	//NewDatabase,
	//NewMessaging,
	NewLogger,
	//NewMetrics,
)

func NewDatabase(config configuration.Configuration) (database.Database, error) {
	return mongodb.NewMongoDatabase(config.Database.ConnectionString, config.Database.ConnectionString)
}

func NewMessaging(config configuration.Configuration) (messaging.Messaging, error) {
	return pulsar.NewPulsarClient(config.Messaging.URL)
}

func NewLogger(config configuration.Configuration) (logging.Logger, error) {
	return zap.NewZapLogger(config.Server.Enviroment)
}

func NewMetrics() (metrics.Metrics, error) {
	return prometheus.NewPrometheusService()
}
