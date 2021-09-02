package server

import (
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func New() (*Server, error) {
	return &Server{
		Router: mux.NewRouter(),
	}, nil
}