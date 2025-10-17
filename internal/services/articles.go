package services

import (
	"database/sql"
	"fmt"

	"github.com/nathanberry97/rss2go/internal/queries"
	"github.com/nathanberry97/rss2go/internal/schema"
)

func GetArticles(conn *sql.DB, page int, limit int) (schema.PaginationResponse, error) {
	if page < 0 || limit <= 0 {
		return schema.PaginationResponse{}, fmt.Errorf("Invalid pagination parameters: page=%d, limit=%d", page, limit)
	}
	offset, nextPage := page*limit, page+1

	query := queries.GetArticlesRecent()
	rows, err := conn.Query(query, limit, offset)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("Failed to execute query: %w", err)
	}
	defer rows.Close()

	articles, err := formatArticles(rows)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("Failed to format articles: %w", err)
	}

	countQuery := queries.GetTotalArticlesRecent()
	var totalItems int
	if err := conn.QueryRow(countQuery).Scan(&totalItems); err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("Failed to get total items: %w", err)
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
		return schema.PaginationResponse{}, fmt.Errorf("Invalid pagination parameters: page=%d, limit=%d", page, limit)
	}
	offset, nextPage := page*limit, page+1

	query := queries.GetArticlesByFeed()

	rows, err := conn.Query(query, id, limit, offset)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("Failed to execute query: %w", err)
	}
	defer rows.Close()

	articles, err := formatArticles(rows)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("Failed to format articles: %w", err)
	}

	countQuery := queries.GetTotalArticlesByFeed()
	var totalItems int
	if err := conn.QueryRow(countQuery, id).Scan(&totalItems); err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("Failed to get total items: %w", err)
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

func InsertArticles(conn *sql.DB, articles []schema.RssItem, feedID int64) error {
	if len(articles) == 0 {
		return nil
	}

	query := queries.InsertArticlesQuery(len(articles))

	args := make([]any, 0, len(articles)*4)
	for _, article := range articles {
		args = append(args, feedID, article.Title, article.Link, article.PubDate)
	}

	_, err := conn.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("Failed to batch insert articles: %w", err)
	}
	return nil
}
