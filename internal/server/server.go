package server

import "github.com/nandanpi/go-mux-jwt-auth/internal/database"

type Server struct {
	listenAddress string
	services      database.DbServices
}

func NewServer(listenAddress string, services database.DbServices) *Server {
	return &Server{
		listenAddress: listenAddress,
		services:      services,
	}
}
