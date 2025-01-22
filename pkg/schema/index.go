package schema

type RssPostBody struct {
	URL string `form:"url" json:"url"`
}

type RssFeed struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	NAME string `json:"name"`
}

type RssItem struct {
	TITLE       string `json:"title"`
	DESCRIPTION string `json:"description"`
	LINK        string `json:"link"`
	PUB_DATE    string `json:"pub_date"`
}

type RssArticle struct {
	FEED_ID int `json:"feed_id"`
	RssItem
}

type RssArticleResponse struct {
	TotalItems int         `json:"total_items"`
	Items      interface{} `json:"items"`
}

type PaginationResponse struct {
	TotalItems int         `json:"total_items"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Items      interface{} `json:"items"`
}
