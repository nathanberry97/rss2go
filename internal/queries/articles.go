package queries

import (
	"fmt"
	"strings"
)

func GetArticlesRecent() string {
	return buildArticleQuery(`a.published_at >= datetime('now', '-31 days')`)
}

func GetArticlesByFeed() string {
	return buildArticleQuery(`f.id = ?`)
}

func GetTotalArticlesRecent() string {
	return buildArticleTotalQuery("articles", `published_at >= datetime('now', '-30 days')`)
}

func GetTotalArticlesByFeed() string {
	return buildArticleTotalQuery("articles", `feed_id = ?`)
}

func InsertArticlesQuery(num int) string {
	placeholders := make([]string, num)

	for i := 0; i < num; i++ {
		placeholders[i] = "(?, ?, ?, ?)"
	}

	return fmt.Sprintf(
		"INSERT OR IGNORE INTO articles (feed_id, title, url, published_at) VALUES %s",
		strings.Join(placeholders, ","),
	)
}
