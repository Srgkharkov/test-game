package apiserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewAPIServer() *APIServer {
	return &APIServer{Router: NewRouter()}
}

type APIServer struct {
	Router *mux.Router
}

func (s *APIServer) Run() error {
	return http.ListenAndServe(":8080", s.Router)
}
