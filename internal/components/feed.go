package components

import (
	"html/template"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateFeedList(feeds []schema.RssFeed) (template.HTML, error) {
	data := schema.RssFeeds{
		Items: feeds,
	}

	return renderTemplates(
		"web/templates/feed/fragments/feeds_list.tmpl",
		"feeds_list",
		data,
	)
}

func GenerateFeedInputForm(endpoint, label string) (template.HTML, error) {
	data := schema.FeedInputForm{
		Endpoint: endpoint,
		Label:    label,
	}

	return renderTemplates(
		"web/templates/feed/fragments/feed_input_form.tmpl",
		"feed_input_form",
		data,
	)
}

func GenerateOPMLButton(endpoint string) (template.HTML, error) {
	data := schema.OPMLFeedBtn{
		Endpoint: endpoint,
	}

	return renderTemplates(
		"web/templates/feed/fragments/opml_buttons.tmpl",
		"opml_buttons",
		data,
	)
}
