package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/nisharyan/golang-rss-aggregator/internal/database"
)

// User is a custom user definition used for responding to API requests. This
// struct is similar to the database User struct but with customised json tags.
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}
