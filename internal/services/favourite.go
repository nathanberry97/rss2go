package services

import (
	"database/sql"
	"fmt"

	"github.com/nathanberry97/rss2go/internal/queries"
	"github.com/nathanberry97/rss2go/internal/schema"
)

func PostFavourite(conn *sql.DB, id string) error {
	query := queries.InsertFavourite()

	_, err := conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to insert article into database: %w", err)
	}

	return nil
}

func DeleteFavourite(conn *sql.DB, id string) error {
	query := queries.DeleteFavourite()

	_, err := conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to insert article into database: %w", err)
	}

	return nil
}

func GetFavourites(conn *sql.DB, page int, limit int) (schema.PaginationResponse, error) {
	if page < 0 || limit <= 0 {
		return schema.PaginationResponse{}, fmt.Errorf("Invalid pagination parameters: page=%d, limit=%d", page, limit)
	}

	offset, nextPage := page*limit, page+1
	query := queries.GetArticlesFavourite()

	rows, err := conn.Query(query, limit, offset)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("Error querying database: %w", err)
	}
	defer rows.Close()

	articles, err := formatArticles(rows)
	if err != nil {
		return schema.PaginationResponse{}, fmt.Errorf("Failed to format articles: %w", err)
	}

	countQuery := queries.GetTotalArticlesFavourite()
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
