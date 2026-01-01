package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT env is not found in configuration file.")
	}

	// Setup the http server with the router.
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create a new Router for v1 version of API.
	v1Router := chi.NewRouter()
	// Register the readiness handler.
	v1Router.Get("/healthz", readinessHandler)
	// Register the err handler.
	v1Router.Get("/error", errorHandler)

	// Mount the v1 router on the main router.
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", portString),
		Handler: router,
	}

	fmt.Printf("Starting server on port %s...\n", portString)
	// Start the server.
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
