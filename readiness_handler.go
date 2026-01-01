package main

import "net/http"

// readinessHandler handles the readiness probe requests.
func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}