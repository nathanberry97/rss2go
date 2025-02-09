package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func InsertArticles(conn *sql.DB, articles []schema.RssItem, feedID int64) error {
	if len(articles) == 0 {
		return nil
	}

	queryPlaceholders := make([]string, 0, len(articles))
	queryArgs := make([]interface{}, 0, len(articles)*4)

	for _, article := range articles {
		queryPlaceholders = append(queryPlaceholders, "(?, ?, ?, ?)")
		queryArgs = append(queryArgs, feedID, article.Title, article.Link, article.PubDate)
	}

	query := fmt.Sprintf(
		"INSERT OR IGNORE INTO articles (feed_id, title, url, published_at) VALUES %s",
		strings.Join(queryPlaceholders, ","),
	)

	_, err := conn.Exec(query, queryArgs...)
	if err != nil {
		return fmt.Errorf("failed to batch insert articles: %w", err)
	}

	return nil
}
