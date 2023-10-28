package apiserver

import (
	"net/http"
	"time"
)

// The Observer method serves as a middleware that measures the request
// processing time and forwards the data to the 'metrics' package.
func (h *APIHandler) Observer(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		inner.ServeHTTP(w, r)

		time := time.Since(start)
		(*h.metrics.RequestDuration).Observe(time.Seconds())
	})
}
