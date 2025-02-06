package services

import (
	"database/sql"
	"fmt"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func insertArticles(conn *sql.DB, articles []schema.RssItem, feedID int64) error {
	stmt, err := conn.Prepare("INSERT OR IGNORE INTO articles (feed_id, title, url, published_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare article insertion statement: %w", err)
	}
	defer stmt.Close()

	for _, article := range articles {
		_, err = stmt.Exec(feedID, article.Title, article.Link, article.PubDate)
		if err != nil {
			fmt.Println("failed to insert article into database:", err)
			return fmt.Errorf("failed to insert article into database: %w", err)
		}
	}

	return nil
}
