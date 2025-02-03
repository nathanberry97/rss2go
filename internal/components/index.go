package components

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateForm(endpoint, label string) template.HTML {
	return template.HTML(`
        <form hx-post="` + endpoint + `" hx-trigger="submit" hx-swap="none" hx-on::after-request="clearInput()">
            <input type="text" id="url" name="url" placeholder="` + label + `" required>
        </form>
    `)
}

func GenerateArticleList(articles schema.PaginationResponse) template.HTML {
	articleItems := ""
	var articlesHTML template.HTML

	for _, article := range articles.Items {
		articleItems += `<li>
			<a href="` + article.Link + `" target="_blank">` + article.Title + `</a>
            <small>` + article.PubDate + `</small>
        </li><br>`
	}

	articlesHTML = template.HTML(`<ul id=articles-list>` + articleItems + `</ul>`)

	if articles.NextPage != -1 {
		articlesHTML += template.HTML(fmt.Sprintf(`
            <div id="articles"
                 hx-trigger="revealed"
                 hx-get="/partials/articles?page=%d"
                 hx-swap="afterend">
            </div>`, articles.NextPage))
	}

	return template.HTML(articlesHTML)
}

func GenerateFeedList(feeds []schema.RssFeed) template.HTML {
	listItems := ""
	for _, feed := range feeds {
		listItems += `<li>
			<a href="` + feed.URL + `" target="_blank">` + feed.Name + `</a>
			<button class="delete-btn"
                    hx-delete="/partials/feed/` + strconv.Itoa(feed.ID) + `"
                    hx-trigger="click"
                    hx-swap="none"
                    data-feed-id="` + strconv.Itoa(feed.ID) + `">
                Delete
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
                <li><a href="/">Articles</a></li>
                <li><a href="/feeds">Feeds</a></li>
              </ul>
            </nav>
        </div>
    `)
}

func GenerateMetaData() template.HTML {
	return template.HTML(`
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>rss2go</title>
        <link rel="stylesheet" href="/static/css/style.css" />
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
        <link
            href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;700|Roboto:wght@400;700&display=swap"
            rel="stylesheet"
        />
        <script src="/static/js/htmx.min.js"></script>
        <script src="/static/js/index.js"></script>
    `)
}
