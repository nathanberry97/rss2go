package services

import (
	"bytes"
	"database/sql"
	"encoding/xml"
	"fmt"
	"strings"
	"sync"

	"github.com/nathanberry97/rss2go/internal/rss"
	"github.com/nathanberry97/rss2go/internal/schema"
)

func PostFeed(conn *sql.DB, postBody schema.RssPostBody) error {
	var name string
	var articles []schema.RssItem

	name, articles, err := rss.FeedHandler(postBody.URL)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error while parsing feed: %w", err)
	}

	query := "INSERT INTO feeds (name, url) VALUES (?, ?)"
	result, err := conn.Exec(query, name, postBody.URL)
	if err != nil {
		return fmt.Errorf("failed to insert feed into database: %w", err)
	}

	feedID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	err = InsertArticles(conn, articles, feedID)
	if err != nil {
		return err
	}

	return nil
}

func GetFeeds(conn *sql.DB) ([]schema.RssFeed, error) {
	query := "SELECT id, name, url FROM feeds"
	rows, err := conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var feeds []schema.RssFeed
	for rows.Next() {
		var feed schema.RssFeed
		if err := rows.Scan(&feed.ID, &feed.Name, &feed.URL); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		feeds = append(feeds, feed)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return feeds, nil
}

func DeleteFeed(conn *sql.DB, id int) error {
	_, err := conn.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return fmt.Errorf("error enabling foreign keys: %w", err)
	}

	query := "DELETE FROM feeds WHERE id = ?"
	_, err = conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting feed: %w", err)
	}

	return nil
}

func PostFeedOpml(conn *sql.DB, opmlData []byte) error {
	var opml schema.OPML
	err := xml.NewDecoder(bytes.NewReader(opmlData)).Decode(&opml)
	if err != nil {
		return fmt.Errorf("failed to parse OPML file: %w", err)
	}

	feedUrls := extractFeedUrls(opml.Body.Outlines)

	var wg sync.WaitGroup
	errChan := make(chan error, len(feedUrls))

	const maxConcurrency = 10
	semaphore := make(chan struct{}, maxConcurrency)

	for _, url := range feedUrls {
		url := strings.TrimSpace(url)
		if url == "" {
			continue
		}

		wg.Add(1)
		go func(feedURL string) {
			defer wg.Done()
			semaphore <- struct{}{}        // acquire slot
			defer func() { <-semaphore }() // release slot

			if err := PostFeed(conn, schema.RssPostBody{URL: feedURL}); err != nil {
				if !strings.Contains(err.Error(), "UNIQUE constraint failed: feeds.url") {
					errChan <- fmt.Errorf("failed to add %s: %w", feedURL, err)
					return
				}
				fmt.Printf("Skipping duplicate feed: %s\n", feedURL)
			}
		}(url)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return err
	}

	fmt.Printf("Successfully imported feeds with articles\n")
	return nil
}

func GetFeedsOpml(conn *sql.DB) ([]byte, error) {
	query := "SELECT name, url FROM feeds"
	rows, err := conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var feeds []schema.OpmlFeed
	for rows.Next() {
		var feed schema.OpmlFeed
		if err := rows.Scan(&feed.Name, &feed.URL); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		feeds = append(feeds, feed)
	}

	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n\n")
	b.WriteString(`<opml version="2.0">` + "\n")
	b.WriteString("  <head>\n")
	b.WriteString("    <title>rss2go Subscriptions</title>\n")
	b.WriteString("  </head>\n")
	b.WriteString("  <body>\n")
	b.WriteString(`    <outline text="rss2go" title="rss2go">` + "\n")

	for _, feed := range feeds {
		feedType := inferFeedType(feed.URL)
		fmt.Fprintf(&b, `      <outline text="%s" type="%s" xmlUrl="%s" />`+"\n",
			xmlEscape(feed.Name), feedType, xmlEscape(feed.URL))
	}

	b.WriteString("    </outline>\n")
	b.WriteString("  </body>\n")
	b.WriteString("</opml>\n")

	return []byte(b.String()), nil
}
