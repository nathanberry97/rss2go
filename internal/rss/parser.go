package rss

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/nathanberry97/rss2go/internal/schema"
)

func FeedHandler(url string) (string, []schema.RssItem, error) {
	var name string
	var articles []schema.RssItem

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return "", nil, err
	}

	name = feed.Title
	for _, item := range feed.Items {
		formattedDate, err := parsePubDate(item.Published)
		if err != nil {
			fmt.Printf("failed to parse pubDate: %v\n", err)
			continue
		}

		article := schema.RssItem{
			Title:   item.Title,
			Link:    item.Link,
			PubDate: formattedDate,
		}

		articles = append(articles, article)
	}

	return name, articles, nil
}

func parsePubDate(pubDate string) (string, error) {
	formats := []string{
		time.DateTime,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, format := range formats {
		formattedDate, err := formatDate(pubDate, format, time.DateTime)
		if err == nil {
			return formattedDate, nil
		}
	}

	return "", fmt.Errorf("Unable to process PubDate")
}

func formatDate(dateStr, inputFormat, outputFormat string) (string, error) {
	t, err := time.Parse(inputFormat, dateStr)
	if err != nil {
		return "", err
	}
	return t.Format(outputFormat), nil
}
