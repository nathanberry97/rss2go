package components

import (
	"fmt"
	"html/template"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateNavbar() template.HTML {
	navbarItems := []schema.NavbarItem{
		{Href: "/", ImgSrc: "/static/images/icons/latest.svg", Alt: "Latest", Label: "Latest"},
		{Href: "/articles/favourites", ImgSrc: "/static/images/icons/favourite.svg", Alt: "Favourites", Label: "Favourites"},
		{Href: "/articles/later", ImgSrc: "/static/images/icons/readlater.svg", Alt: "Read Later", Label: "Read Later"},
		{Href: "/feeds", ImgSrc: "/static/images/favicon.svg", Alt: "Feeds", Label: "Feeds"},
		{Href: "/settings", ImgSrc: "/static/images/icons/settings.svg", Alt: "Settings", Label: "Settings"},
	}

	navbarHTML := `<div class="navbar">
        <div class="navbar__logo">
            rss<span class="navbar__logo-number">2</span>go
        </div>
        <nav class="navbar__navigation">
            <ul class="navbar__list">`

	for _, item := range navbarItems {
		navbarHTML += fmt.Sprintf(`
            <li class="navbar__item">
                <a class="navbar__link" href="%s">
                    <img class="navbar__icon" src="%s" alt="%s">
                    %s
                </a>
            </li>`, item.Href, item.ImgSrc, item.Alt, item.Label)
	}

	navbarHTML += `
            </ul>
        </nav>
    </div>`

	return template.HTML(navbarHTML)
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
