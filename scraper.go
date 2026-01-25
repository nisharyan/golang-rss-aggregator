package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nisharyan/golang-rss-aggregator/internal/database"
)

func scrapeFeeds(
	concurreny int,
	timeBetweenRequest time.Duration,
	db *database.Queries,
) {
	log.Println("Starting feed scraper with concurrency:", concurreny, "and interval:", timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurreny))
		if err != nil {
			log.Println("Error fetching feeds to scrape:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, feed, wg)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	// Mark the feeds as fetched.
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", feed.Url, "error:", err)
		return
	}

	rssFeed, err := urlToRSSFeed(feed.Url)
	if err != nil {
		log.Println("Error scraping feed:", feed.Url, "error:", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println("Error parsing pubDate for item:", item.Link, "error:", err)
			continue
		}
		_, err = db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Url:         item.Link,
				FeedID:      feed.ID,
				Description: description,
				PublishedAt: pubDate,
			},
		)
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key") {
				log.Println("Error creating post:", item.Link, "error:", err)
			}
			continue
		}
	}
	log.Printf("Scraped feed: %s with %d items\n", rssFeed.Channel.Title, len(rssFeed.Channel.Items))
}
