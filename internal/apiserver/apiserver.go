package apiserver

import (
	"github.com/Srgkharkov/test-game/internal/game"
	"github.com/Srgkharkov/test-game/internal/metrics"
	"github.com/gorilla/mux"
	"net/http"
)

func NewAPIServer(game *game.Game, metrics *metrics.Metrics) *APIServer {
	return &APIServer{Router: NewRouter(game, metrics)}
}

type APIServer struct {
	Router *mux.Router
}

func (s *APIServer) Run(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}
