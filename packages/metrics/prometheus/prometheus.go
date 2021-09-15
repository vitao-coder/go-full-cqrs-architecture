package prometheus

import (
	"github.com/vitao-coder/go-full-cqrs-architecture/packages/metrics/model"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusService struct {
	httpRequestHistogram *prometheus.HistogramVec
}

func NewPrometheusService() (*PrometheusService, error) {
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &PrometheusService{
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

func (s *PrometheusService) SaveHTTP(appMetrics *model.AplicationMetrics) {
	s.httpRequestHistogram.WithLabelValues(appMetrics.Handler, appMetrics.Method, appMetrics.StatusCode).Observe(appMetrics.Duration)
}
