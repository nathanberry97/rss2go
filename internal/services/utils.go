package services

import (
	"bytes"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"strings"
	"sync"
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

		article.Title = html.UnescapeString(article.Title)

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
	case duration < 24*time.Hour:
		hours := int(duration.Hours())
		if hours <= 1 {
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

func parseOpml(data []byte) (*schema.OPML, error) {
	var opml schema.OPML
	if err := xml.NewDecoder(bytes.NewReader(data)).Decode(&opml); err != nil {
		return nil, fmt.Errorf("failed to parse OPML file: %w", err)
	}
	return &opml, nil
}

func importFeedsConcurrently(conn *sql.DB, feedUrls []string) error {
	const maxConcurrency = 10

	var wg sync.WaitGroup
	errChan := make(chan error, len(feedUrls))
	semaphore := make(chan struct{}, maxConcurrency)

	for _, u := range feedUrls {
		url := strings.TrimSpace(u)
		if url == "" {
			continue
		}

		wg.Add(1)
		go func(feedURL string) {
			defer wg.Done()
			semaphore <- struct{}{}        // acquire slot
			defer func() { <-semaphore }() // release slot

			if err := importSingleFeed(conn, feedURL); err != nil {
				errChan <- err
			}
		}(url)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return err
	}

	return nil
}

func importSingleFeed(conn *sql.DB, feedURL string) error {
	err := PostFeed(conn, schema.RssPostBody{URL: feedURL})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: feeds.url") {
			fmt.Printf("Skipping duplicate feed: %s\n", feedURL)
			return nil
		}
		return fmt.Errorf("failed to add %s: %w", feedURL, err)
	}
	return nil
}
