package queries

func GetArticlesReadLater() string {
	return buildArticleQuery(`a.id IN (SELECT article_id FROM read_later)`)
}

func GetTotalArticlesReadLater() string {
	return buildArticleTotalQuery("read_later")
}

func InsertReadLater() string {
	return "INSERT INTO read_later (article_id) VALUES (?)"
}

func DeleteReadLater() string {
	return "DELETE FROM read_later WHERE article_id = ?"
}
