package rss

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/nathanberry97/rss2go/internal/schema"
	"github.com/nathanberry97/rss2go/internal/utils"
)

func PostFeedHandler(url string) (string, []schema.RssItem, error) {
	err := checkValidFeed(url)
	if err != nil {
		return "", nil, err
	}

	name, err := parseRssTitle(url)
	if err != nil {
		return "", nil, err
	}

	articles, err := parseRssItems(url)
	if err != nil {
		return "", nil, err
	}

	return name, articles, nil
}

func checkValidFeed(url string) error {
	isValid := false

	c := colly.NewCollector()

	c.OnXML("//rss", func(e *colly.XMLElement) {
		isValid = true
	})

	if err := c.Visit(url); err != nil {
		return fmt.Errorf("failed to visit URL %q: %w", url, err)
	}

	if isValid == false {
		return fmt.Errorf("not a valid rss feed")
	}

	return nil
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

func parseRssItems(url string) ([]schema.RssItem, error) {
	var articles []schema.RssItem

	c := colly.NewCollector()

	c.OnXML("//rss/channel/item", func(e *colly.XMLElement) {

		pubDate := e.ChildText("pubDate")
		formattedDate, err := parsePubDate(pubDate)
		if err != nil {
			fmt.Printf("failed to parse pubDate: %v\n", err)
			return
		}

		article := schema.RssItem{
			Title:   e.ChildText("title"),
			Link:    e.ChildText("link"),
			PubDate: formattedDate,
		}

		articles = append(articles, article)
	})

	if err := c.Visit(url); err != nil {
		return nil, fmt.Errorf("failed to visit URL %q: %w", url, err)
	}

	return articles, nil
}

func parsePubDate(pubDate string) (string, error) {
	formats := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
	}

	for _, format := range formats {
		formattedDate, err := utils.FormatDate(pubDate, format, time.DateTime)
		if err == nil {
			return formattedDate, nil
		}
	}

	return "", fmt.Errorf("Unable to process pubDate")
}
