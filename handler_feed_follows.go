package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/go-chi/chi"
	"github.com/nisharyan/golang-rss-aggregator/internal/database"
)

// createFeedFollow handles the creation of a new fedd_follow.
// This is handler for an authenticated endpoint.
func (apiConfig *apiConfig) createFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	// Define a struct to parse the JSON request body and store the parameters
	// needed to create a new user.
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	// Parse the JSON request body.
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON request body: %v", err))
		return
	}

	// Insert the new feed follow into the database.
	feed_follow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	// Respond with the custom feed follow definition which is similar to as
	// the database feed follow definition but with customised json tags.
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

// getFeedFollowsHandler handles the retrieval of all feed follows for a
// specific user. This is an authenticated endpoint.
func (apiConfig *apiConfig) getFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error retrieving feed follows: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feed_follows))
}

// deleteFeedFollowHandler deletes a feed follow for a specific user.
// This is an authenticated endpoint.
func (apiConfig *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feed_follow_id")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid feed_follow_id: %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error retrieving feed follows: %v", err))
		return
	}

	respondWithJSON(w, 200, map[string]string{"message": "Feed follow deleted successfully"})
}