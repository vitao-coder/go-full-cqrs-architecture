package metrics

import "github.com/vitao-coder/go-full-cqrs-architecture/packages/metrics/model"

type Metrics interface {
	SaveHTTP(appMetrics *model.AplicationMetrics)
}
