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
		return s.handleCreateAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// function to return a user by ID
func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := GetID(r)
		if err != nil {
			return err
		}

		account, err := s.db.GetUserByID(id)
		if err != nil {
			return err
		}

		WriteJSON(w, http.StatusOK, account)
	}

	return fmt.Errorf("Method %v not allowed ", r.Method)
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
	account, err := NewAccount(req.UserName, req.Email, req.Password)
	if err != nil {
		return err
	}

	if err := s.db.CreateUserAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)

}
