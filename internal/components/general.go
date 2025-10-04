package components

import (
	"html/template"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateNavbar() (template.HTML, error) {
	data := []schema.NavbarItem{
		{Href: "/", Label: "Latest"},
		{Href: "/articles/favourites", Label: "Favourites"},
		{Href: "/articles/later", Label: "Read Later"},
		{Href: "/feeds", Label: "Feeds"},
	}

	return renderTemplates(
		"web/templates/general/navbar.tmpl",
		"navbar",
		data,
	)
}

func GenerateMetaData(cssFile string) (template.HTML, error) {
	data := schema.Metadata{CssFile: cssFile}

	return renderTemplates(
		"web/templates/general/metadata.tmpl",
		"metadata",
		data,
	)
}
