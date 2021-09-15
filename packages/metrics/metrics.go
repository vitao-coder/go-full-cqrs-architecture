package metrics

import "go-full-cqrs-architecture/packages/metrics/model"

type Metrics interface {
	SaveApplicationMetrics(appMetrics *model.AplicationMetrics)
}
