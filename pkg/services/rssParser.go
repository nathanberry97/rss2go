package services

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/nathanberry97/rss2go/pkg/schema"
)

func parseRssItems(url string) ([]schema.RssItem, error) {
	var articles []schema.RssItem

	c := colly.NewCollector()

	c.OnXML("//rss/channel/item", func(e *colly.XMLElement) {
		pubDate := e.ChildText("pubDate")

		// Try parsing the first format (with timezone offset)
		parsedDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", pubDate)
		if err != nil {
			// If the first format fails, try parsing the second format (with GMT)
			parsedDate, err = time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", pubDate)
			if err != nil {
				fmt.Printf("failed to parse pubDate: %v\n", err)
				return
			}
		}

		// Format the date into the desired format (YYYY-MM-DD HH:MM:SS)
		formattedDate := parsedDate.Format("2006-01-02 15:04:05")

		article := schema.RssItem{
			TITLE:       e.ChildText("title"),
			DESCRIPTION: e.ChildText("description"),
			LINK:        e.ChildText("link"),
			PUB_DATE:    formattedDate, // Use the formatted date
		}

		articles = append(articles, article)
	})

	if err := c.Visit(url); err != nil {
		return nil, fmt.Errorf("failed to visit URL %q: %w", url, err)
	}

	return articles, nil
}

func parseRssTitle(url string) (string, error) {
	var name string

	c := colly.NewCollector()

	c.OnXML("//rss/channel/title", func(e *colly.XMLElement) {
		name = e.Text
	})

	if err := c.Visit(url); err != nil {
		return "", fmt.Errorf("failed to visit URL %q: %w", url, err)
	}

	if name == "" {
		return "", fmt.Errorf("unable to extract title from RSS feed")
	}

	return name, nil
}
