package main

type Server struct {
	listenAddress string
	services      dbServices
}

func NewServer(listenAddress string, services dbServices) *Server {
	return &Server{
		listenAddress: listenAddress,
		services:      services,
	}
}
