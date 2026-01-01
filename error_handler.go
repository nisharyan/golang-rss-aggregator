package main

import "net/http"

// error handler to respond with error messages.
func errorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "This is the default error message.")
}