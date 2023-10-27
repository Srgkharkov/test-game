package apiserver

import (
	"github.com/Srgkharkov/test-game/internal/game"
	"github.com/gorilla/mux"
	"net/http"
)

func NewAPIServer(g *game.Game) *APIServer {
	return &APIServer{Router: NewRouter(g)}
}

type APIServer struct {
	Router *mux.Router
}

func (s *APIServer) Run() error {
	return http.ListenAndServe(":8080", s.Router)
}
