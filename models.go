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
	APIKey    string    `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

// Feed is a custom user definition that is created for the same reason as the
// User above.
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	customFeeds := make([]Feed, len(dbFeeds))
	for idx, dbFeed := range dbFeeds {
		customFeeds[idx] = databaseFeedtoFeed(dbFeed)
	}
	return customFeeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowtoFeedFollow(dbFeed database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		UserID:    dbFeed.UserID,
		FeedID:    dbFeed.FeedID,
	}
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	customFeedFollows := make([]FeedFollow, len(dbFeedFollows))
	for idx, dbFeedFollow := range dbFeedFollows {
		customFeedFollows[idx] = databaseFeedFollowToFeedFollow(dbFeedFollow)
	}

	return customFeedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
}

func databasePostToPost(dbPost database.Post) Post {
	description := (*string)(nil)
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
		Description: description,
	}
}

func databasePostsToPost(dbPosts []database.Post) []Post {
	customPosts := make([]Post, len(dbPosts))
	for idx, dbPost := range dbPosts{
		customPosts[idx] = databasePostToPost(dbPost)
	}
	return customPosts
}
