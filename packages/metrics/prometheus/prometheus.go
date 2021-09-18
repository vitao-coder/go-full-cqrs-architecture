package prometheus

import (
	"fmt"
	"strconv"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/logging"
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/metrics/model"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusService struct {
	httpRequestHistogram *prometheus.HistogramVec
	logger               logging.Logger
}

func NewPrometheusService(logger logging.Logger) (*PrometheusService, error) {
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &PrometheusService{
		httpRequestHistogram: http,
		logger:               logger,
	}
	err := prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

func (s *PrometheusService) SaveHTTP(appMetrics model.AplicationMetrics) {
	statusStrCode := strconv.Itoa(appMetrics.StatusCode)
	s.httpRequestHistogram.WithLabelValues(appMetrics.Handler, appMetrics.Method, statusStrCode).Observe(appMetrics.Duration.Seconds())
	s.logger.Debug(fmt.Sprintf("Metrics - Handler: %s - Method: %s - StatusCode: %s - Took %s", appMetrics.Handler, appMetrics.Method, statusStrCode, appMetrics.Duration))
}
