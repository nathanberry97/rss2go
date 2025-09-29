package queries

func GetFeeds() string {
	return "SELECT id, name, url FROM feeds ORDER BY name"
}

func GetFeedsOpml() string {
	return "SELECT name, url FROM feeds"
}

func InsertFeed() string {
	return "INSERT INTO feeds (name, url) VALUES (?, ?)"
}

func DeleteFeed() string {
	return "DELETE FROM feeds WHERE id = ?"
}
