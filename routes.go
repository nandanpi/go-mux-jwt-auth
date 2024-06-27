package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/", JWTAuthMiddleware(makeHandlerFunc(s.HandleRootRoute))).Methods("GET")
	r.HandleFunc("/signup", makeHandlerFunc(s.HandleSignUp)).Methods("POST")
	r.HandleFunc("/login", makeHandlerFunc(s.HandleLogin)).Methods("POST")
	r.HandleFunc("/allaccounts", makeHandlerFunc(s.HandleGetAllAccounts)).Methods("GET")

	log.Println("Starting server on: ", s.listenAddress)

	http.ListenAndServe(s.listenAddress, r)
}
