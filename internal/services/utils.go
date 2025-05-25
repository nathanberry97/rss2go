package services

import (
	"database/sql"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/nathanberry97/rss2go/internal/schema"
	"github.com/nathanberry97/rss2go/internal/utils"
)

func GetArticleName(conn *sql.DB, id string) (string, error) {
	query := "SELECT name FROM feeds WHERE id = ?"

	var name string
	if err := conn.QueryRow(query, id).Scan(&name); err != nil {
		return "", fmt.Errorf("failed to get feed name for feed id: %w", err)
	}

	return name, nil
}

func formatArticles(rows *sql.Rows) ([]schema.RssArticle, error) {
	var articles []schema.RssArticle

	for rows.Next() {
		var article schema.RssArticle
		if err := rows.Scan(&article.FeedName, &article.FeedId, &article.Id, &article.Title, &article.Link, &article.PubDate, &article.Fav, &article.Later); err != nil {
			return []schema.RssArticle{}, fmt.Errorf("failed to scan row: %w", err)
		}

		formatPubDate, err := utils.FormatTime(article.PubDate, time.RFC3339)
		if err != nil {
			return []schema.RssArticle{}, fmt.Errorf("failed parsing pubDate: %w", err)
		}

		article.PubDate = formatPubDate
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return []schema.RssArticle{}, fmt.Errorf("rows iteration error: %w", err)
	}

	return articles, nil
}

func extractFeedUrls(outlines []schema.OpmlOutline) []string {
	var urls []string
	for _, outline := range outlines {
		if outline.XMLURL != "" {
			urls = append(urls, outline.XMLURL)
		}
		if len(outline.Outlines) > 0 {
			urls = append(urls, extractFeedUrls(outline.Outlines)...)
		}
	}
	return urls
}

func xmlEscape(s string) string {
	return html.EscapeString(s)
}

func inferFeedType(url string) string {
	if strings.Contains(url, ".atom") || strings.Contains(url, "/atom") || strings.Contains(url, "format=atom") {
		return "atom"
	}
	return "rss"
}
