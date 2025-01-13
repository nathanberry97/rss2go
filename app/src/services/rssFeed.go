package services

import (
	"database/sql"
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/nathanberry97/rss2go/src/schema"
)

func PostRssFeed(conn *sql.DB, postBody schema.RssPostBody) (int64, error) {
	var name string

	c := colly.NewCollector()
	c.OnXML("//rss/channel/title", func(e *colly.XMLElement) {
		name = e.Text
	})

	err := c.Visit(postBody.URL)
	if err != nil {
		fmt.Printf("Error visiting URL: %v\n", err)
		return 0, err
	}

	if name == "" {
		fmt.Println("Error: Unable to extract title from RSS feed")
		return 0, fmt.Errorf("unable to extract title from RSS feed")
	}

	query := "INSERT INTO feeds (name, url) VALUES (?, ?)"
	result, err := conn.Exec(query, name, postBody.URL)
	if err != nil {
		fmt.Printf("Error inserting feed into database: %v\n", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Error retrieving last insert ID: %v\n", err)
		return 0, err
	}

	fmt.Printf("RSS feed inserted successfully with ID: %d\n", id)
	return id, nil
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
		if err := rows.Scan(&feed.ID, &feed.NAME, &feed.URL); err != nil {
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
	query := "DELETE FROM feeds WHERE id = ?"
	_, err := conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting feed: %w", err)
	}

	return nil
}
