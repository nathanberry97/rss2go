package schema

type RssPostBody struct {
	URL string `json:"url"`
}

type RssFeed struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	NAME string `json:"name"`
}
