package main

import "net/http"

type HandlerType func(w http.ResponseWriter, r *http.Request) error

type ServerError struct {
	error string
}

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
