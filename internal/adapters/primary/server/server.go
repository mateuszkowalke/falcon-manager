package server

import (
	"net/http"

	"falcon-manager/internal/core/services"
)

type Server struct {
	mux http.Handler
}

func InitServer(accService *services.AccountService) (*Server, error) {
	mux := initRoutes(accService)
	return &Server{mux}, nil
}

func (srv *Server) Run() error {
	return http.ListenAndServe(":3000", srv.mux)
}
