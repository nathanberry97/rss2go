package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/nathanberry97/rss2go/internal/components"
	"github.com/nathanberry97/rss2go/internal/queries"
	"github.com/nathanberry97/rss2go/internal/rss"
	"github.com/nathanberry97/rss2go/internal/schema"
)

func GetFeeds(conn *sql.DB) ([]schema.RssFeed, error) {
	query := queries.GetFeeds()
	rows, err := conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error querying database: %w", err)
	}
	defer rows.Close()

	var feeds []schema.RssFeed
	for rows.Next() {
		var feed schema.RssFeed
		if err := rows.Scan(&feed.ID, &feed.Name, &feed.URL); err != nil {
			return nil, fmt.Errorf("Error scanning row: %w", err)
		}
		feeds = append(feeds, feed)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error during rows iteration: %w", err)
	}

	return feeds, nil
}

func PostFeed(conn *sql.DB, postBody schema.RssPostBody) error {
	var name string
	var articles []schema.RssItem

	name, articles, err := rss.FeedHandler(strings.TrimSpace(postBody.URL))
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error while parsing feed: %w", err)
	}

	query := queries.InsertFeed()
	result, err := conn.Exec(query, name, postBody.URL)
	if err != nil {
		return fmt.Errorf("Failed to insert feed into database: %w", err)
	}

	feedID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Failed to retrieve last insert ID: %w", err)
	}

	err = InsertArticles(conn, articles, feedID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFeed(conn *sql.DB, id int) error {
	_, err := conn.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return fmt.Errorf("Error enabling foreign keys: %w", err)
	}

	query := queries.DeleteFeed()
	_, err = conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error deleting feed: %w", err)
	}

	return nil
}

func PostFeedOpml(conn *sql.DB, opmlData []byte) error {
	opml, err := parseOpml(opmlData)
	if err != nil {
		return err
	}

	feedUrls := extractFeedUrls(opml.Body.Outlines)
	if len(feedUrls) == 0 {
		return fmt.Errorf("No feed URLs found in OPML file")
	}

	if err := importFeedsConcurrently(conn, feedUrls); err != nil {
		return err
	}

	fmt.Println("Successfully imported feeds with articles")
	return nil
}

func GetFeedsOpml(conn *sql.DB) ([]byte, error) {
	query := queries.GetFeedsOpml()
	rows, err := conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error querying database: %w", err)
	}
	defer rows.Close()

	var feeds []schema.OpmlFeed
	for rows.Next() {
		var feed schema.OpmlFeed
		if err := rows.Scan(&feed.Name, &feed.URL); err != nil {
			return nil, fmt.Errorf("Error scanning row: %w", err)
		}
		feed.Type = inferFeedType(feed.URL)
		feeds = append(feeds, feed)
	}

	return components.RenderRSSTemplate(
		"web/templates/feed/fragments/opml.tmpl",
		"feeds_opml",
		feeds,
	)
}
