package model

import "time"

type AplicationMetrics struct {
	Handler    string
	Method     string
	StatusCode string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewAplicationMetrics(handler string, method string) *AplicationMetrics {
	return &AplicationMetrics{
		Handler: handler,
		Method:  method,
	}
}

func (a *AplicationMetrics) Started() {
	a.StartedAt = time.Now()
}

func (a *AplicationMetrics) Finished() {
	a.FinishedAt = time.Now()
	a.Duration = time.Since(a.StartedAt).Seconds()
}
