package components

import (
	"fmt"
	"html/template"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateArticleQuery(query schema.QueryKey, feedId *string) (template.HTML, error) {
	var path string

	switch query {
	case schema.Articles:
		path = "/partials/articles?page=0"
	case schema.ArticlesFavourite:
		path = "/partials/favourite?page=0"
	case schema.ArticlesReadLater:
		path = "/partials/later?page=0"
	case schema.ArticlesByFeed:
		path = fmt.Sprintf("/partials/articles/%s?page=0", *feedId)
	}

	data := schema.ArticleQuery{
		Path: path,
	}

	return renderTemplates(
		"web/templates/articles/fragments/article_query.tmpl",
		"article_query",
		data,
	)
}

func GenerateArticleList(articles schema.PaginationResponse, feedId *int, query schema.QueryKey) (template.HTML, error) {
	var nextPageURL string
	if articles.NextPage != -1 {
		switch query {
		case schema.Articles:
			nextPageURL = fmt.Sprintf(`/partials/articles?page=%d`, articles.NextPage)
		case schema.ArticlesFavourite:
			nextPageURL = fmt.Sprintf(`/partials/favourite?page=%d`, articles.NextPage)
		case schema.ArticlesReadLater:
			nextPageURL = fmt.Sprintf(`/partials/later?page=%d`, articles.NextPage)
		case schema.ArticlesByFeed:
			nextPageURL = fmt.Sprintf(`/partials/articles/%d?page=%d`, *feedId, articles.NextPage)
		}
	}

	data := struct {
		Items       []schema.RssArticle
		NextPage    int
		NextPageURL string
	}{
		Items:       articles.Items,
		NextPage:    articles.NextPage,
		NextPageURL: nextPageURL,
	}

	return renderTemplates(
		"web/templates/articles/fragments/articles_list.tmpl",
		"articles_list",
		data,
	)
}

func GenerateArticleButton(path, name, target string, del bool) (template.HTML, error) {
	data := schema.ArticleBtn{
		Path:   path,
		Name:   name,
		Target: target,
		Delete: del,
	}

	return renderTemplates(
		"web/templates/articles/fragments/article_button.tmpl",
		"article_button",
		data,
	)
}
