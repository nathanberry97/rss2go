package services

import (
	"database/sql"
	"fmt"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func PostRssFeed(conn *sql.DB, postBody schema.RssPostBody) (int64, error) {
	var name string
	var articles []schema.RssItem

	err := checkValidFeed(postBody.URL)
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("Not a valid RSS feed: %w", err)
	}

	articles, err = parseRssItems(postBody.URL)
	if err != nil {
		return 0, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	name, err = parseRssTitle(postBody.URL)
	if err != nil {
		return 0, fmt.Errorf("failed to parse RSS feed title: %w", err)
	}

	query := "INSERT INTO feeds (name, url) VALUES (?, ?)"
	result, err := conn.Exec(query, name, postBody.URL)
	if err != nil {
		return 0, fmt.Errorf("failed to insert feed into database: %w", err)
	}

	feedID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	stmt, err := conn.Prepare("INSERT INTO articles (feed_id, title, description, url, published_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare article insertion statement: %w", err)
	}
	defer stmt.Close()

	for _, article := range articles {
		_, err = stmt.Exec(feedID, article.Title, article.Description, article.Link, article.PubDate)
		if err != nil {
			fmt.Println("failed to insert article into database:", err)
			return 0, fmt.Errorf("failed to insert article into database: %w", err)
		}
	}

	return feedID, nil
}

func GetRssFeeds(conn *sql.DB) ([]schema.RssFeed, error) {
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

func DeleteRssFeed(conn *sql.DB, id int) error {
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
