package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	if statusCode >= 500 {
		log.Printf("Internal server error: %s", message)
	}

	// Define a struct to hold the error message.
	// The 'json' tags specify the key names in the JSON response. When the
	// errorResponse struct is marshaled into JSON, the field 'Error' will
	// appear as 'error' in the resulting JSON object.
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, statusCode, errorResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	// Marshal the payload into JSON format.
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON request : %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Add a header field to include the json in the response.
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
