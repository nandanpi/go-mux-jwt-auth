package types

import "net/http"

type HandlerType func(w http.ResponseWriter, r *http.Request) error

type ServerError struct {
	Error string
}

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
