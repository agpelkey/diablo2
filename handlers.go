package main

import (
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

 