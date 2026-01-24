package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nisharyan/golang-rss-aggregator/internal/database"
)

// createFeedHandler handles the creation of a new feed.
func (apiConfig *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	// Define a struct to parse the JSON request body and store the parameters
	// needed to create a new user.
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	// Parse the JSON request body.
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON request body: %v", err))
		return
	}

	// Insert the new feed into the database.
	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	// Respond with the custom feed definition which is similar to as the
	// database user definition but with customised json tags.
	respondWithJSON(w, 201, databaseFeedtoFeed(feed))
}

// getFeedsHandler handles the retrieval of all feeds. This is an
// unauthenticated endpoint.
func (apiConfig *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error retrieving feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedstoFeeds(feeds))
}
