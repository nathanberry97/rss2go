package services

import (
	"database/sql"
	"fmt"

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
