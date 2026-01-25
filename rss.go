package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
	"fmt"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToRSSFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a custom request instead of using httpClient.Get()
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RSSFeed{}, err
    }

    // Set a browser-like User-Agent
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := httpClient.Do(req)
	if err != nil {
		return RSSFeed{}, err
	}
	defer resp.Body.Close()

	//Check the status code
    if resp.StatusCode != http.StatusOK {
        return RSSFeed{}, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
    }

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)

	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
