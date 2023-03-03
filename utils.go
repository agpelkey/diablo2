package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func GetID(r *http.Request) (int, error) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("Invalid id %v is given: ", id)
	}
	return id, nil
}
