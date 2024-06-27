package utils

import (
	"encoding/json"
	"net/http"

	"github.com/nandanpi/go-mux-jwt-auth/internal/types"
)

func MakeHandlerFunc(f types.HandlerType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, types.ServerError{Error: "Something went wrong"})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WritePlainText(w http.ResponseWriter, status int, v string) error {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
