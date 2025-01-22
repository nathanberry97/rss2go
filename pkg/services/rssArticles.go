package services

import (
	"database/sql"
	"fmt"

	"github.com/nathanberry97/rss2go/pkg/schema"
)

func GetRssArticles(conn *sql.DB, page int, limit int) (schema.PaginationResponse, error) {
	if page < 0 || limit <= 0 {
		return schema.PaginationResponse{}, fmt.Errorf("invalid pagination parameters: page=%d, limit=%d", page, limit)
	}

	offset := page * limit

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
		if err := rows.Scan(&article.FEED_ID, &article.TITLE, &article.DESCRIPTION, &article.LINK, &article.PUB_DATE); err != nil {
			return schema.PaginationResponse{}, fmt.Errorf("failed to scan row: %w", err)
		}
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

	response := schema.PaginationResponse{
		TotalItems: totalItems,
		Page:       page,
		Limit:      limit,
		Items:      articles,
	}

	return response, nil
}
