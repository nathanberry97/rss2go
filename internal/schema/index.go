package schema

// Data models
type RssItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	PubDate     string `json:"pub_date"`
}

type RssArticle struct {
	FeedId int `json:"feed_id"`
	RssItem
}

// Request bodies
type RssPostBody struct {
	URL string `form:"url" json:"url"`
}

// Response Structures
type RssFeed struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
}

type RssArticleResponse struct {
	TotalItems int         `json:"total_items"`
	Items      interface{} `json:"items"`
}

type PaginationResponse struct {
	TotalItems int `json:"total_items"`
	NextPage   int `json:"page"`
	Limit      int `json:"limit"`
	Items      []RssArticle
}
