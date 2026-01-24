package main

import (
	"net/http"
	"github.com/nisharyan/golang-rss-aggregator/internal/auth"
	"github.com/nisharyan/golang-rss-aggregator/internal/database"
	"fmt"
)

// The authedHandler is a function type that represents an HTTP handler
// which requires an authenticated user. All functions to retriever user,
// user feeds all require an authenticated user.
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// middlewareAuth is a middleware function that wraps any authedHandler
// function which requires user authentication before proceeding to the actual
// handler logic. The middlewareAuth function takes an authedHandler as input
// and returns an http.HandlerFunc that can be used as a standard HTTP handler.
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Error retrieving API key: %v", err))
			return
		}

		// The Context is a built-in Go feature that carries deadlines,
		// cancellation signals, and other request-scoped values across
		// API boundaries and between processes.
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), api_key)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}