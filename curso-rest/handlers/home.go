package handlers

import (
	"encoding/json"
	"net/http"

	"jjgdevelopment.com/go/rest-ws/server"
)

type NewResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(NewResponse{
			Message: "Welcome to Platzi Go",
			Status:  "true",
		})
	}
}
