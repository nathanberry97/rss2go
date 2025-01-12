package services

import (
	"context"
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/jackc/pgx/v5"
	"github.com/nathanberry97/rss2go/src/schema"
)

func PostRssFeed(conn *pgx.Conn, postBody schema.RssPostBody) (int, error) {
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
		return 0, err
	}

	query := "INSERT INTO rss.feeds (name, url) VALUES ($1, $2) RETURNING id"
	var id int
	err = conn.QueryRow(context.Background(), query, name, postBody.URL).Scan(&id)
	if err != nil {
		fmt.Printf("Error inserting feed into database: %v\n", err)
		return 0, err
	}

	fmt.Printf("RSS feed inserted successfully with ID: %d\n", id)
	return id, nil
}

func GetRssFeeds(conn *pgx.Conn) ([]schema.RssFeed, error) {
	query := "SELECT id, name, url FROM rss.feeds"
	rows, err := conn.Query(context.Background(), query)
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
