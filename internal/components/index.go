package components

import (
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

func GenerateArticleList(articles []schema.RssArticle) template.HTML {
	articleItems := ""

	for _, article := range articles {
		articleItems += `<li>
			<a href="` + article.Link + `" target="_blank">` + article.Title + `</a>
            <br>
            <small>` + article.PubDate + `</small>
        </li>`
	}

	return template.HTML(`<ul>` + articleItems + `</ul>`)
}

func GenerateFeedList(feeds []schema.RssFeed) template.HTML {
	listItems := ""
	for _, feed := range feeds {
		listItems += `<li>
			<a href="` + feed.URL + `" target="_blank">` + feed.Name + `</a>
			<button class="delete-btn"
                    hx-delete="/rss_feed/` + strconv.Itoa(feed.ID) + `"
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
                <li style="margin-right: 20px;"><a href="/" style="text-decoration: none;">Home</a></li>
                <li><a href="/feeds" style="text-decoration: none;">Feeds</a></li>
              </ul>
            </nav>
        </div>
    `)
}
