package queries

func GetArticlesFavourite() string {
	return buildArticleQuery(`a.id IN (SELECT article_id FROM favourites)`)
}

func GetTotalArticlesFavourite() string {
	return buildArticleTotalQuery("favourites")
}

func InsertFavourite() string {
	return "INSERT INTO favourites (article_id) VALUES (?)"
}

func DeleteFavourite() string {
	return "DELETE FROM favourites WHERE article_id = ?"
}
