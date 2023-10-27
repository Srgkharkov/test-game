package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	RequestsTotal      *prometheus.Counter
	ErrorResponseTotal *prometheus.Counter
	RequestDuration    *prometheus.Histogram
}

func NewMetrics() *Metrics {
	requestsTotal := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
	)
	errorResponseTotal := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_responses_with_error_total",
			Help: "Total number of HTTP responses with error.",
		},
	)
	requestDuration := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests.",
		},
	)

	return &Metrics{
		RequestsTotal:      &requestsTotal,
		ErrorResponseTotal: &errorResponseTotal,
		RequestDuration:    &requestDuration,
	}
}

func (m *Metrics) Run() {
	prometheus.MustRegister(*m.RequestsTotal, *m.ErrorResponseTotal, *m.RequestDuration)
}
