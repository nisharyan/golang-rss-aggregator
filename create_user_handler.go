package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nisharyan/golang-rss-aggregator/internal/database"
)

// readinessHandler handles the readiness probe requests.
// The *apiConfig struct is passed to access the database if needed.
// The * in the receiver indicates that this method has a pointer receiver.
// The pointer receiver allows the method to modify the apiConfig struct if
// needed.
func (apiConfig *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct to parse the JSON request body and store the parameters
	// needed to create a new user.
	type parameters struct {
		Name string `json:"name"`
	}

	// Parse the JSON request body.
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON request body: %v", err))
		return
	}

	// Insert the new user into the database.
	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	// Respond with the custom user definition which is similar to as the
	// database user definition but with customised json tags.
	respondWithJSON(w, 200, databaseUsertoUser(user))
}
