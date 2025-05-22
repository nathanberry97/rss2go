package components

import (
	"fmt"
	"html/template"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func GenerateNavbar() template.HTML {
	navbarItems := []schema.NavbarItem{
		{Href: "/", Label: "Latest"},
		{Href: "/articles/favourites", Label: "Favourites"},
		{Href: "/articles/later", Label: "Read Later"},
		{Href: "/feeds", Label: "Feeds"},
	}

	navbarHTML := `<div class="navbar">
		<div class="navbar__header">
			<div class="navbar__logo">
				rss<span class="navbar__logo-number">2</span>go
			</div>
			<button class="navbar__hamburger" aria-label="Toggle menu">
				<img src="/static/images/icons/menu.svg" class="navbar__hamburger-icon" alt="Menu">
			</button>
		</div>
        <nav class="navbar__navigation">
            <ul class="navbar__list">`

	for _, item := range navbarItems {
		navbarHTML += fmt.Sprintf(`
            <li class="navbar__item">
                <a class="navbar__link" href="%s">
                    %s
                </a>
            </li>`, item.Href, item.Label)
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
