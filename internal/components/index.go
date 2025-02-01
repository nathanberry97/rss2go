package components

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateForm(endpoint, label string) template.HTML {
	return template.HTML(`
        <form hx-post="` + endpoint + `" hx-trigger="submit" hx-swap="none">
            <label for="url">` + label + `</label>
            <input type="text" id="url" name="url" required>
            <button type="submit">Submit</button>
        </form>
    `)
}

func GenerateArticleList(articles schema.PaginationResponse) template.HTML {
	articleItems := ""
	var articlesHTML template.HTML

	for _, article := range articles.Items {
		articleItems += `<li>
			<a href="` + article.Link + `" target="_blank">` + article.Title + `</a>
            <br>
            <small>` + article.PubDate + `</small>
        </li>`
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
        <div class="navbar-container">
            <nav class="navbar-nav">
              <ul style="list-style-type: none; padding: 0; margin: 0; display: flex;">
                <li style="margin-right: 20px;"><a href="/" style="text-decoration: none;">Articles</a></li>
                <li><a href="/feeds" style="text-decoration: none;">Feeds</a></li>
              </ul>
            </nav>
        </div>
    `)
}
