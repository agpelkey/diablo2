package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
	db         PostgresRepo
}

// function to create new APIServer
func NewAPIServer(listenAddr string, db PostgresRepo) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

// function signature we are using in this app
type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

// decorate our apiFunc into an http handlerfun
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()

	//mux.HandleFunc("/", makeHTTPHandleFunc(ReturnHome))

	log.Println("API Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, mux)
}
