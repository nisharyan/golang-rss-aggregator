package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/nisharyan/golang-rss-aggregator/internal/database"

	// Include the Postgres driver. The '_' indicates that we are importing
	// the package solely for its side-effects (i.e., its init function).
	// This pacakge is required by the sql package to interact with Postgres
	// databases.
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT env is not found in configuration file.")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL env is not found in configuration file.")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	apiConfig := &apiConfig{
		DB: database.New(conn),
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
	// Register the create user handler.
	v1Router.Post("/users", apiConfig.createUserHandler)

	// Mount the v1 router on the main router.
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", portString),
		Handler: router,
	}

	fmt.Printf("Starting server on port %s...\n", portString)
	// Start the server.
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
