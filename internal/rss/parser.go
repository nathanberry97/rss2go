package rss

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/nathanberry97/rss2go/internal/schema"
	"github.com/nathanberry97/rss2go/internal/utils"
)

func FeedHandler(url string) (string, []schema.RssItem, error) {
	var name string
	var articles []schema.RssItem

	feedType, err := checkValidFeed(url)
	if err != nil {
		return "", nil, err
	}

	name, err = parseFeedTitle(url, feedType)
	if err != nil {
		return "", nil, err
	}

	articles, err = parseFeedItems(url, feedType)
	if err != nil {
		return "", nil, err
	}

	return name, articles, nil
}

func checkValidFeed(url string) (schema.FeedType, error) {
	feedType := schema.FeedTypeNone
	c := colly.NewCollector()

	c.OnXML("//rss", func(e *colly.XMLElement) {
		feedType = schema.FeedTypeRSS
	})

	c.OnXML("//feed", func(e *colly.XMLElement) {
		feedType = schema.FeedTypeAtom
	})

	if err := c.Visit(url); err != nil {
		return schema.FeedTypeNone, fmt.Errorf("failed to visit URL %q: %w", url, err)
	}

	if feedType == schema.FeedTypeNone {
		return schema.FeedTypeNone, fmt.Errorf("not a valid rss or atom feed")
	}

	return feedType, nil
}

func parseFeedTitle(url string, feedType schema.FeedType) (string, error) {
	var title, name string
	c := colly.NewCollector()

	switch feedType {
	case schema.FeedTypeRSS:
		title = "//rss/channel/title"
	case schema.FeedTypeAtom:
		title = "//feed/title"
	default:
		return "", fmt.Errorf("Unsupported feed type: %v", feedType)
	}

	c.OnXML(title, func(e *colly.XMLElement) {
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

func parseFeedItems(url string, feedType schema.FeedType) ([]schema.RssItem, error) {
	var articles []schema.RssItem
	var item, dateField string
	c := colly.NewCollector()

	switch feedType {
	case schema.FeedTypeRSS:
		item = "//rss/channel/item"
		dateField = "pubDate"
	case schema.FeedTypeAtom:
		item = "//feed/entry"
		dateField = "updated"
	default:
		return nil, fmt.Errorf("Unsupported feed type: %v", feedType)
	}

	c.OnXML(item, func(e *colly.XMLElement) {

		pubDate := e.ChildText(dateField)
		formattedDate, err := parsePubDate(pubDate)
		if err != nil {
			fmt.Printf("failed to parse pubDate: %v\n", err)
			return
		}

		link := e.ChildAttr("link", "href")
		if link == "" {
			link = e.ChildText("link")
		}

		article := schema.RssItem{
			Title:   e.ChildText("title"),
			Link:    link,
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
