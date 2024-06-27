package server

import (
	"encoding/json"
	"net/http"

	"github.com/nandanpi/go-mux-jwt-auth/internal/auth"
	"github.com/nandanpi/go-mux-jwt-auth/internal/types"
	"github.com/nandanpi/go-mux-jwt-auth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleRootRoute(w http.ResponseWriter, r *http.Request) error {
	return utils.WritePlainText(w, http.StatusOK, "Root Route")
}

func (s *Server) HandleSignUp(w http.ResponseWriter, r *http.Request) error {
	accReq := &types.CreateAccountRequest{}

	err := json.NewDecoder(r.Body).Decode(accReq)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	s.services.CreateAccount(accReq.Username, string(hashedPassword))

	return utils.WritePlainText(w, http.StatusOK, "Account created")
}

func (s *Server) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	loginReq := &types.LoginRequest{}
	json.NewDecoder(r.Body).Decode(loginReq)

	dbUser := s.services.GetUser(w, loginReq.Username)

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginReq.Password))
	if err != nil {
		return utils.WriteJSON(w, http.StatusUnauthorized, types.ServerError{Error: "Invalid creds"})
	}

	token, err := auth.GenerateJWT(dbUser.ID)
	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, types.ServerError{Error: "Couldn't generate JWT Token"})
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
