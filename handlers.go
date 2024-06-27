package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleRootRoute(w http.ResponseWriter, r *http.Request) error {
	return WritePlainText(w, http.StatusOK, "Root Route")
}

func (s *Server) HandleSignUp(w http.ResponseWriter, r *http.Request) error {
	accReq := &CreateAccountRequest{}

	err := json.NewDecoder(r.Body).Decode(accReq)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	s.services.createAccount(accReq.Username, string(hashedPassword))

	return WritePlainText(w, http.StatusOK, "Account created")
}

func (s *Server) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	loginReq := &LoginRequest{}
	json.NewDecoder(r.Body).Decode(loginReq)

	dbUser := s.services.getUser(w, loginReq.Username)

	log.Println(dbUser)

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginReq.Password))
	if err != nil {
		return WriteJSON(w, http.StatusUnauthorized, ServerError{error: "Invalid creds"})
	}

	token, err := GenerateJWT(dbUser.ID)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ServerError{error: "Couldn't generate JWT Token"})
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (s *Server) HandleGetAllAccounts(w http.ResponseWriter, r *http.Request) error {
	users, err := s.services.getAllUsers(w)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, users)
}
