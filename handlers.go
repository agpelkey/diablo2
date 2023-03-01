package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
func ReturnHome(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Hello World")
	return nil
}
*/

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		// logic for getting an account
	}

	if r.Method == "POST" {
		// handle logic for creating accounts
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// function to handle account creation
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	// populate variable with Account request struct
	req := new(CreateAccountRequest)

	// decore request body
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	// populate account variable with fields from req
	account, err := NewAccount(req.FirstName, req.LastName, req.UserName, req.Email, req.Password)
	if err != nil {
		return err
	}

	if err := s.db.CreateUserAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)

}
