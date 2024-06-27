package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandanpi/go-mux-jwt-auth/internal/auth"
	"github.com/nandanpi/go-mux-jwt-auth/internal/utils"
)

func (s *Server) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/", auth.JWTAuthMiddleware(utils.MakeHandlerFunc(s.HandleRootRoute))).Methods("GET")
	r.HandleFunc("/signup", utils.MakeHandlerFunc(s.HandleSignUp)).Methods("POST")
	r.HandleFunc("/login", utils.MakeHandlerFunc(s.HandleLogin)).Methods("POST")

	log.Println("Starting server on: ", s.listenAddress)

	http.ListenAndServe(s.listenAddress, r)
}
