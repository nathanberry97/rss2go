package schema

import (
	"database/sql"
	"encoding/xml"
)

type QueryKey string

const (
	Articles          QueryKey = "articles"
	ArticlesReadLater QueryKey = "articlesReadLater"
	ArticlesFavourite QueryKey = "articlesFavourite"
	ArticlesByFeed    QueryKey = "articlesByFeed"
)

// Data models
type OpmlOutline struct {
	XMLName  xml.Name      `xml:"outline"`
	Text     string        `xml:"text,attr"`
	Title    string        `xml:"title,attr"`
	Type     string        `xml:"type,attr"`
	XMLURL   string        `xml:"xmlUrl,attr"`
	Outlines []OpmlOutline `xml:"outline"`
}

type OPML struct {
	Body struct {
		Outlines []OpmlOutline `xml:"outline"`
	} `xml:"body"`
}

type RssItem struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Link    string `json:"link"`
	PubDate string `json:"pub_date"`
	Fav     bool   `json:"is_fav"`
	Later   bool   `json:"is_read_later"`
}

type Task struct {
	FeedId int64
	URL    string
	Conn   *sql.DB
}

type RssArticle struct {
	FeedId   int    `json:"feed_id"`
	FeedName string `json:"feed_name"`
	RssItem
}

type NavbarItem struct {
	Href  string
	Label string
}

// Request bodies
type RssPostBody struct {
	URL string `form:"url" json:"url"`
}

// Response Structures
type OpmlFeed struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type RssFeed struct {
	ID int `json:"id"`
	OpmlFeed
}

type RssArticleResponse struct {
	TotalItems int `json:"total_items"`
	Items      any `json:"items"`
}

type PaginationResponse struct {
	TotalItems int `json:"total_items"`
	NextPage   int `json:"page"`
	Limit      int `json:"limit"`
	Items      []RssArticle
}
