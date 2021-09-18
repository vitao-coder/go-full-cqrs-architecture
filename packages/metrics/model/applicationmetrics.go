package model

import "time"

type AplicationMetrics struct {
	Handler    string
	Method     string
	StatusCode int
	StartedAt  time.Time
	Duration   time.Duration
}

func NewAplicationMetrics(handler string, method string, statusCode int, startedAt time.Time, duration time.Duration) *AplicationMetrics {
	return &AplicationMetrics{
		Handler:    handler,
		Method:     method,
		StatusCode: statusCode,
		StartedAt:  startedAt,
		Duration:   duration,
	}
}
