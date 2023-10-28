package apiserver

import (
	"github.com/Srgkharkov/test-game/internal/game"
	"github.com/Srgkharkov/test-game/internal/metrics"
	"github.com/gorilla/mux"
	"net/http"
)

// NewAPIServer initializes the router and provides the necessary repository to the handlers.
func NewAPIServer(game *game.Game, metrics *metrics.Metrics) *APIServer {
	return &APIServer{Router: NewRouter(game, metrics)}
}

// APIServer structure includes mux.Router
type APIServer struct {
	Router *mux.Router
}

// The Run method takes an address and port, such as ":8080," launches the server on that port, and delegates routing.
func (s *APIServer) Run(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}
