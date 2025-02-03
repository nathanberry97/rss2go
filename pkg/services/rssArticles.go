package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GetRssArticles(conn *sql.DB, page int, limit int) (schema.PaginationResponse, error) {
	if page < 0 || limit <= 0 {
		return schema.PaginationResponse{}, fmt.Errorf("invalid pagination parameters: page=%d, limit=%d", page, limit)
	}

	offset := page * limit
	nextPage := page + 1

	query := `
        SELECT id, title, description, url, published_at
        FROM articles
        ORDER BY published_at DESC
        LIMIT ? OFFSET ?
    `

	rows, err := conn.Query(query, limit, offset)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var articles []schema.RssArticle
	for rows.Next() {
		var article schema.RssArticle
		if err := rows.Scan(&article.FeedId, &article.Title, &article.Description, &article.Link, &article.PubDate); err != nil {
			return schema.PaginationResponse{}, fmt.Errorf("failed to scan row: %w", err)
		}

		t, err := time.Parse(time.RFC3339, article.PubDate)
		if err != nil {
			return schema.PaginationResponse{}, fmt.Errorf("failed parsing pubDate: %w", err)
		}

		formatPubDate := t.Format("01/02/2006 15:04")
		article.PubDate = formatPubDate

		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("rows iteration error: %w", err)
	}

	countQuery := "SELECT COUNT(*) FROM articles"
	var totalItems int
	if err := conn.QueryRow(countQuery).Scan(&totalItems); err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("failed to get total items: %w", err)
	}

	remainingItems := totalItems - (page * limit)
	if remainingItems <= 0 {
		nextPage = -1
	}

	response := schema.PaginationResponse{
		TotalItems: totalItems,
		NextPage:   nextPage,
		Limit:      limit,
		Items:      articles,
	}

	return response, nil
}
