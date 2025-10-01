package components

import (
	"fmt"
	"html/template"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateArticleQuery(query schema.QueryKey, feedId *string) template.HTML {
	var path string

	switch query {
	case schema.Articles:
		path = "/partials/articles?page=0"
	case schema.ArticlesFavourite:
		path = "partials/favourite?page=0"
	case schema.ArticlesReadLater:
		path = "/partials/later?page=0"
	case schema.ArticlesByFeed:
		path = fmt.Sprintf("partials/articles/%s?page=0", *feedId)
	}

	return template.HTML(fmt.Sprintf(`
		<div id="articles"
			 hx-trigger="revealed"
			 hx-get="%s"
			 hx-swap="afterend">
		</div>
	`, path))
}

func GenerateArticleList(articles schema.PaginationResponse, feedId *int, query schema.QueryKey) template.HTML {
	articleItems := ""
	var articlesHTML template.HTML
	var tarFav, tarLater = "favourite", "readlater"

	for _, article := range articles.Items {
		articleItems += `<li class="articles__item">
			<a class="articles__link articles__details-text articles__details-text--article" href="` + article.Link + `" target="_blank">` + article.Title + `</a>
			<div id="` + tarFav + `_` + article.Id + `" class="articles__btn-container">
				` + GenerateArticleButton(`/partials/favourite/`+article.Id, "Favourite", tarFav+`_`+article.Id, article.Fav) + `
			</div>
			<div id="` + tarLater + `_` + article.Id + `" class="articles__btn-container">
				` + GenerateArticleButton(`/partials/later/`+article.Id, "Read Later", tarLater+`_`+article.Id, article.Later) + `
			</div>
			<span class="articles__details-text articles__details-text--date">` + article.PubDate + `</span>
		</li>`
	}

	articlesHTML = template.HTML(`<ul class="articles__list">` + articleItems + `</ul>`)

	if articles.NextPage != -1 {
		var nextPageURL string

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

		articlesHTML += template.HTML(fmt.Sprintf(`
            <div id="articles"
                 hx-trigger="revealed"
                 hx-get="%s"
                 hx-swap="afterend">
            </div>`, nextPageURL))
	}

	return template.HTML(articlesHTML)
}

func GenerateArticleButton(path, name, target string, del bool) string {
	class, trigger := "articles__btn articles__btn--post", "hx-post"
	if del {
		class, trigger = "articles__btn articles__btn--delete", "hx-delete"
	}

	return fmt.Sprintf(`
        <button class="%s"
            %s="%s"
            hx-target="#%s"
            hx-swap="innerHTML">
            %s
        </button>
    `, class, trigger, path, target, name)
}
