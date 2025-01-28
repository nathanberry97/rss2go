package components

import (
	"github.com/nathanberry97/rss2go/pkg/schema"
	"html/template"
)

func GenerateForm(endpoint, label string) template.HTML {
	return template.HTML(`
        <form hx-post="/` + endpoint + `" hx-target="#feedList" hx-swap="innerHTML"
            <label for="url">` + label + `</label>
            <input type="text" id="url" name="url" required>
            <button type="submit">Submit</button>
        </form>
    `)
}

func GenerateFeedList(feeds []schema.RssFeed) template.HTML {
	listItems := ""
	for _, feed := range feeds {
		listItems += `<li><a href="` + feed.URL + `" target="_blank">` + feed.NAME + `</a></li>`
	}
	return template.HTML(`<ul>` + listItems + `</ul>`)
}
