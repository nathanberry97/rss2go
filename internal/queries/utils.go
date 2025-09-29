package queries

import (
	"fmt"
	"strings"
)

func buildArticleQuery(conditions ...string) string {
	query := `
		SELECT f.name, a.feed_id, a.id, a.title, a.url, a.published_at,
			EXISTS (SELECT 1 FROM favourites fav WHERE fav.article_id = a.id) AS is_fav,
			EXISTS (SELECT 1 FROM read_later rl WHERE rl.article_id = a.id) AS is_read_later
		FROM articles a
		JOIN feeds f ON a.feed_id = f.id
	`

	if len(conditions) > 0 {
		query += "\nWHERE " + strings.Join(conditions, " AND ")
	}

	query += "\n" + `
		ORDER BY published_at DESC
		LIMIT ? OFFSET ?
	`

	return query
}

func buildArticleTotalQuery(table string, conditions ...string) string {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)

	if len(conditions) > 0 {
		query += "\nWHERE " + strings.Join(conditions, " AND ")
	}

	return query
}
