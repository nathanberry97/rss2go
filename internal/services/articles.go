package services

import (
	"database/sql"
	"fmt"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GetArticles(conn *sql.DB, page int, limit int) (schema.PaginationResponse, error) {
	if page < 0 || limit <= 0 {
		return schema.PaginationResponse{}, fmt.Errorf("invalid pagination parameters: page=%d, limit=%d", page, limit)
	}
	offset, nextPage := page*limit, page+1

	query := `
        SELECT f.name, a.feed_id, a.id, a.title, a.url, a.published_at,
            EXISTS (SELECT 1 FROM favourites fav WHERE fav.article_id = a.id) AS is_fav,
            EXISTS (SELECT 1 FROM read_later rl WHERE rl.article_id = a.id) AS is_read_later
        FROM articles a
        JOIN feeds f ON a.feed_id = f.id
		WHERE a.published_at >= datetime('now', '-31 days')
        ORDER BY published_at DESC
        LIMIT ? OFFSET ?
    `
	rows, err := conn.Query(query, limit, offset)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	articles, err := formatArticles(rows)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("failed to format articles: %w", err)
	}

	countQuery := "SELECT COUNT(*) FROM articles WHERE published_at >= datetime('now', '-30 days')"
	var totalItems int
	if err := conn.QueryRow(countQuery).Scan(&totalItems); err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("failed to get total items: %w", err)
	}

	remainingItems := totalItems - ((page + 1) * limit)
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

func GetArticlesByFeedId(conn *sql.DB, page int, limit int, id int) (schema.PaginationResponse, error) {
	if page < 0 || limit <= 0 {
		return schema.PaginationResponse{}, fmt.Errorf("invalid pagination parameters: page=%d, limit=%d", page, limit)
	}
	offset, nextPage := page*limit, page+1

	query := `
        SELECT f.name, a.feed_id, a.id, a.title, a.url, a.published_at,
            EXISTS (SELECT 1 FROM favourites fav WHERE fav.article_id = a.id) AS is_fav,
            EXISTS (SELECT 1 FROM read_later rl WHERE rl.article_id = a.id) AS is_read_later
        FROM articles a
        JOIN feeds f ON a.feed_id = f.id
        WHERE f.id = ?
        ORDER BY published_at DESC
        LIMIT ? OFFSET ?
    `

	rows, err := conn.Query(query, id, limit, offset)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	articles, err := formatArticles(rows)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("failed to format articles: %w", err)
	}

	countQuery := "SELECT COUNT(*) FROM articles WHERE feed_id = ?"
	var totalItems int
	if err := conn.QueryRow(countQuery, id).Scan(&totalItems); err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("failed to get total items: %w", err)
	}

	remainingItems := totalItems - ((page + 1) * limit)
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
