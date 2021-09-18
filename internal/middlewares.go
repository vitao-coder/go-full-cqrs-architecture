package server

import (
	"net/http"
	"time"

	httpInternal "github.com/vitao-coder/go-full-cqrs-architecture/internal/http"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/metrics"

	"github.com/vitao-coder/go-full-cqrs-architecture/packages/metrics/model"
)

func MetricsMiddleware(metrics metrics.Metrics) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			pathURL := r.URL.Path
			method := r.Method
			startedAt := time.Now()
			rwWrapped := httpInternal.NewResponseWriter(w)
			next.ServeHTTP(rwWrapped, r)
			finishedAt := time.Since(startedAt)
			appMetric := model.NewAplicationMetrics(pathURL, method, rwWrapped.Status(), startedAt, finishedAt)
			metrics.SaveHTTP(*appMetric)
		}
		return http.HandlerFunc(fn)
	}
}
