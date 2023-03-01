package main

import (
	"log"
	"net/http"
)

// API server struct
type APIServer struct {
	listenAddr string
	db         PostgresRepo
}

// API error struct
type APIError struct {
	Error string `json:"error"`
}

// fucntion that handles the routing as well as starting the server.
func (s *APIServer) Run() {
	mux := http.NewServeMux()

	//mux.HandleFunc("/", makeHTTPHandleFunc(ReturnHome))
	mux.HandleFunc("/acount", makeHTTPHandleFunc(s.handleAccount))

	log.Println("API Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, mux)
}

// function to create new APIServer
func NewAPIServer(listenAddr string, db PostgresRepo) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

// decorate our apiFunc into an http handlerfun
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
