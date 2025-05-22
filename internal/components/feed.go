package components

import (
	"html/template"
	"strconv"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateFeedInputForm(endpoint, label string) template.HTML {
	return template.HTML(`
        <form class="feed__form" hx-post="` + endpoint + `" hx-trigger="submit" hx-swap="none" hx-on::after-request="clearInput()">
            <div class="feed__form-group">
                <input class="feed__input" type="text" id="url" name="url" placeholder="` + label + `" required>
                <button class="feed__submit" type="submit">
                    <img class="feed__submit-icon" src="/static/images/icons/search.svg" alt="Search">
                </button>
            </div>
        </form>
    `)
}

func GenerateFeedList(feeds []schema.RssFeed) template.HTML {
	listItems := ""
	for _, feed := range feeds {
		listItems += `<li class="feed__item">
			<a class="feed__link" href="/articles/` + strconv.Itoa(feed.ID) + `">` + feed.Name + `</a>
			<button class="feed__delete-btn"
                    hx-delete="/partials/feed/` + strconv.Itoa(feed.ID) + `"
                    hx-trigger="click"
                    hx-swap="none"
                    data-feed-id="` + strconv.Itoa(feed.ID) + `">
                <img class="feed__delete-icon" src="/static/images/icons/delete.svg" alt="Delete">
            </button>
		</li>`
	}
	return template.HTML(`<ul class="feed__list">` + listItems + `</ul>`)
}
