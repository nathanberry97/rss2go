# rss2go

![rss2go](./.assets/rss2go.png)

> rss2go is a self-hosted RSS reader built with Go

- [Link to docker hub](https://hub.docker.com/repository/docker/nathanberry97/rss2go/general)

## Getting started

If you want to run the application locally ensure you have `sass` and `go`
installed then run the following command:

```
make run
```

Which will allow you to access the website on `localhost:8080`.

Or if you have `podman` installed you can run the following command:

```
make container
```

## Overview

Folder structure:

```txt
.
├── cmd
│   └── app
├── internal
│   ├── components
│   ├── database
│   ├── routes
│   ├── rss
│   ├── schema
│   ├── scss
│   ├── services
│   ├── utils
│   └── worker
└── web
    ├── scss
    │   ├── abstracts
    │   ├── base
    │   ├── components
    │   ├── layout
    │   ├── main.scss
    │   └── pages
    ├── static
    │   ├── images
    │   └── js
    └── templates
```

If you want to see the current make commands run `make` in your terminal, but
here is a list of the current make commands:

```txt
rss2go

Usage:
  setup                 Install pre-commit hooks
  build                 Build rss2go api
  run                   Build and run rss2go api
  test                  Test backend for rss2go app
  clean                 Clean up build artifacts
  container             Run a local containised version of the application
```

> *Note* to use the container command you will need podman install

This project uses the following technology:

- [go](https://go.dev/)
  - [gin](https://github.com/gin-gonic/gin)
  - [gofeed](https://github.com/mmcdole/gofeed)
  - [sqlite](https://pkg.go.dev/modernc.org/sqlite)
- [sass](https://sass-lang.com/)
- [htmx](https://htmx.org/)
- JavaScript

## OPML

The application supports OPML to bulk upload your feeds, to find this feature
head over to the feeds page and upload your XML file.
There is an example one in the repo called `opml.xml` or you can download
the following to save into an XML file to upload:

```txt
<?xml version="1.0" encoding="UTF-8"?>

<opml version="2.0">
  <head>
    <title>rss2go Subscriptions</title>
  </head>
  <body>
    <outline text="rss2go" title="rss2go">

      <outline text="WIRED" type="rss" xmlUrl="https://www.wired.com/feed/rss" />

      <outline text="Linux.com" type="rss" xmlUrl="https://www.linux.com/feed/" />

      <outline text="Coding Horror" type="rss" xmlUrl="https://blog.codinghorror.com/rss" />

      <outline text="BBC News" type="rss" xmlUrl="http://feeds.bbci.co.uk/news/technology/rss.xml" />

      <outline text="Technology in government" type="rss" xmlUrl="https://technology.blog.gov.uk/feed/" />

      <outline text="Docker" type="rss" xmlUrl="https://www.docker.com/blog/feed/" />

      <outline text="Stack Overflow Blog" type="rss" xmlUrl="https://stackoverflow.blog/feed/" />

      <outline text="Framework Blog" type="rss" xmlUrl="https://frame.work/gb/en/blog.rss" />

      <outline text="The GitHub Blog" type="rss" xmlUrl="https://github.blog/feed/" />

      <outline text="The Verge" type="rss" xmlUrl="https://www.theverge.com/rss/index.xml" />

      <outline text="David Heinemeier Hansson" type="atom" xmlUrl="https://world.hey.com/dhh/feed.atom" />

      <outline text="Hacker News: Front Page" type="rss" xmlUrl="https://hnrss.org/frontpage" />

      <outline text="AWS News Blog" type="rss" xmlUrl="https://aws.amazon.com/blogs/aws/feed/" />

      <outline text="The Go Blog" type="atom" xmlUrl="http://blog.golang.org/feed.atom" />

      <outline text="Nathan Berry" type="rss" xmlUrl="https://nathanberry.co.uk/feed.xml" />

      <outline text="Diomidis D. Spinellis Web Log" type="rss" xmlUrl="https://www.spinellis.gr/blog/dds-blog-rss.xml" />

      <outline text="Dan Bartlett: coach, writer, engineer &amp; founder" type="rss" xmlUrl="https://danbartlett.co.uk/index.xml" />

      <outline text="Jason Fried" type="atom" xmlUrl="https://world.hey.com/jason/feed.atom" />

      <outline text="NASA" type="rss" xmlUrl="https://www.nasa.gov/feed/" />

      <outline text="Kris Tun" type="rss" xmlUrl="https://kristun.dev/rss.xml" />

    </outline>
  </body>
</opml>
```

## API Endpoints

```txt
# HTML Pages

GET    /                           # Articles homepage
GET    /feeds                      # View all feeds
GET    /articles/favourites        # Favourite articles
GET    /articles/later             # Read-later articles
GET    /articles/{feedId}          # Articles by feed ID

# HTMX Partials

GET    /partials/articles          # All articles (partial)
GET    /partials/articles/{id}     # All feed articles (partial)
GET    /partials/favourite         # Favourite list (partial)
GET    /partials/later             # Read-later list (partial)
GET    /partials/feed              # Feed form (partial)
GET    /partials/feed/opml         # OPML import form

POST   /partials/favourite/{id}    # Mark article as favourite
POST   /partials/later/{id}        # Save article for later
POST   /partials/feed              # Add new feed
POST   /partials/feed/opml         # Upload OPML file

DELETE /partials/favourite/{id}    # Remove from favourite
DELETE /partials/later/{id}        # Remove from read-later
DELETE /partials/feed/{id}         # Delete feed
```
