package components

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateArticleQuery(query schema.QueryKey, feedId *string) template.HTML {
	var queryTemplate template.HTML

	switch query {
	case schema.Articles:
		queryTemplate = `
        <div id="articles"
             hx-trigger="revealed"
             hx-get="/partials/articles?page=0"
             hx-swap="afterend">
        </div>
        `
	case schema.ArticlesFavourite:
		queryTemplate = `
        <div id="articles"
             hx-trigger="revealed"
             hx-get="/partials/favourite?page=0"
             hx-swap="afterend">
        </div>
        `
	case schema.ArticlesReadLater:
		queryTemplate = `
        <div id="articles"
             hx-trigger="revealed"
             hx-get="/partials/later?page=0"
             hx-swap="afterend">
        </div>
        `
	case schema.ArticlesByFeed:
		queryTemplate = template.HTML(fmt.Sprintf(`
        <div id="articles"
             hx-trigger="revealed"
             hx-get="/partials/articles/%s?page=0"
             hx-swap="afterend">
        </div>
        `, *feedId))
	}

	return template.HTML(queryTemplate)
}

func GenerateInputForm(endpoint, label string) template.HTML {
	return template.HTML(`
        <form hx-post="` + endpoint + `" hx-trigger="submit" hx-swap="none" hx-on::after-request="clearInput()">
            <div class="input-wrapper">
                <input type="text" id="url" name="url" placeholder="` + label + `" required>
                <button type="submit">
                    <img src="/static/images/icons/search.svg" alt="Search" width="20" height="20" class="search-icon">
                </button>
            </div>
        </form>
    `)
}

func GenerateArticleList(articles schema.PaginationResponse, feedId *int, query schema.QueryKey) template.HTML {
	articleItems := ""
	var articlesHTML template.HTML
	var tarFav, tarLater = "favourite", "readlater"

	for _, article := range articles.Items {
		articleItems += `<li>
			<a href="` + article.Link + `" target="_blank">` + article.Title + `</a>
            <div class="details">
                <small>` + article.FeedName + `</small>
                <div class="buttons-container">
                    <div id="` + tarFav + `_` + article.Id + `">
                        ` + GenerateButton(`/partials/favourite/`+article.Id, "Favourite", tarFav+`_`+article.Id, article.Fav) + `
                    </div>
                    <div id="` + tarLater + `_` + article.Id + `">
                        ` + GenerateButton(`/partials/later/`+article.Id, "Read Later", tarLater+`_`+article.Id, article.Later) + `
                    </div>
                </div>
                <small>` + article.PubDate + `</small>
            </div>
        </li><br>`
	}

	articlesHTML = template.HTML(`<ul id=articles-list>` + articleItems + `</ul>`)

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

func GenerateFeedList(feeds []schema.RssFeed) template.HTML {
	listItems := ""
	for _, feed := range feeds {
		listItems += `<li>
			<a href="/articles/` + strconv.Itoa(feed.ID) + `">` + feed.Name + `</a>
			<button class="delete-btn"
                    hx-delete="/partials/feed/` + strconv.Itoa(feed.ID) + `"
                    hx-trigger="click"
                    hx-swap="none"
                    data-feed-id="` + strconv.Itoa(feed.ID) + `">
                <img src="/static/images/icons/delete.svg" alt="Delete" width="20" height="20" class="delete-icon">
            </button>
		</li>`
	}
	return template.HTML(`<ul>` + listItems + `</ul>`)
}

func GenerateNavbar() template.HTML {
	return template.HTML(`
        <div class="navbar">
            <div class="navbar-logo">
                <span class="logo-text">rss</span><span class="logo-number">2</span><span class="logo-text">go</span>
            </div>
            <nav class="navbar-nav">
              <ul>
                <li>
                    <a href="/">
                        <img src="/static/images/icons/latest.svg" alt="Latest" width="20" height="20" class="latest-icon">
                        Latest
                    </a>
                </li>
                <li>
                    <a href="/articles/favourites">
                        <img src="/static/images/icons/favourite.svg" alt="Favourites" width="20" height="20" class="feed-icon">
                        Favourites
                    </a>
                </li>
                <li>
                    <a href="/articles/later">
                        <img src="/static/images/icons/readlater.svg" alt="Read Later" width="20" height="20" class="feed-icon">
                        Read Later
                    </a>
                </li>
                <li>
                    <a href="/feeds">
                        <img src="/static/images/favicon.svg" alt="Feeds" width="20" height="20" class="feed-icon">
                        Feeds
                    </a>
                </li>
                <li>
                    <a href="/settings">
                        <img src="/static/images/icons/settings.svg" alt="Settings" width="20" height="20" class="feed-icon">
                        Settings
                    </a>
                </li>
              </ul>
            </nav>
        </div>
    `)
}

func GenerateMetaData(cssFile string) template.HTML {
	return template.HTML(`
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>rss2go</title>
        <link rel="icon" type="image/svg+xml" href="/static/images/favicon.svg" />
        <link rel="stylesheet" href="/static/css/` + cssFile + `" />
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
        <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;700&display=swap" rel="stylesheet" />
        <script src="/static/js/htmx.min.js"></script>
        <script src="/static/js/index.js"></script>
    `)
}

func GenerateButton(path, name, target string, del bool) string {
	class, trigger := "post", "hx-post"
	if del {
		class, trigger = "delete", "hx-delete"
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
