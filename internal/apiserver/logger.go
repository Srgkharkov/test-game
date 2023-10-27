package apiserver

import (
	"log"
	"net/http"
	"time"
)

//func Logger(inner http.Handler, name string) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//
//		inner.ServeHTTP(w, r)
//
//		//header := w.Header().
//		//header = header
//
//		log.Printf(
//			"%s %s %s %s",
//			r.Method,
//			r.RequestURI,
//			name,
//			time.Since(start),
//		)
//	})
//}

func (h *APIHandler) Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		inner.ServeHTTP(w, r)

		time := time.Since(start)
		(*h.metrics.RequestDuration).Observe(time.Seconds())
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time,
		)
	})
}
