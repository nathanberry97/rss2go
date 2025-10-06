package services

import (
	"database/sql"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/nathanberry97/rss2go/internal/schema"
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

		formatPubDate, err := formatTime(article.PubDate, time.RFC3339)
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

func formatTime(dateStr, inputFormat string) (string, error) {
	t, err := time.Parse(inputFormat, dateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse time: %w", err)
	}

	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return "just now", nil
	case duration < time.Hour:
		mins := int(duration.Minutes())
		if mins == 1 {
			return "1 minute ago", nil
		}
		return fmt.Sprintf("%d minutes ago", mins), nil
	case duration < 24*time.Hour:
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago", nil
		}
		return fmt.Sprintf("%d hours ago", hours), nil
	case duration < 30*24*time.Hour:
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago", nil
		}
		return fmt.Sprintf("%d days ago", days), nil
	case duration < 365*24*time.Hour:
		months := int(duration.Hours() / (24 * 30))
		if months == 1 {
			return "1 month ago", nil
		}
		return fmt.Sprintf("%d months ago", months), nil
	default:
		years := int(duration.Hours() / (24 * 365))
		if years == 1 {
			return "1 year ago", nil
		}
		return fmt.Sprintf("%d years ago", years), nil
	}
}
