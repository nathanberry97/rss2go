package components

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateNavbar() (template.HTML, error) {
	navbarItems := []schema.NavbarItem{
		{Href: "/", Label: "Latest"},
		{Href: "/articles/favourites", Label: "Favourites"},
		{Href: "/articles/later", Label: "Read Later"},
		{Href: "/feeds", Label: "Feeds"},
	}

	tmpl := template.Must(template.ParseFiles("web/templates/general/navbar.tmpl"))

	var sb strings.Builder
	err := tmpl.Execute(&sb, navbarItems)
	if err != nil {
		return "", fmt.Errorf("Failed to execute navbar template: %v", err)
	}

	return template.HTML(sb.String()), nil
}

func GenerateMetaData(cssFile string) (template.HTML, error) {
	data := schema.Metadata{CssFile: cssFile}
	tmpl := template.Must(template.ParseFiles("web/templates/general/metadata.tmpl"))

	var sb strings.Builder
	err := tmpl.Execute(&sb, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute metadata template: %v", err)
	}

	return template.HTML(sb.String()), nil
}
