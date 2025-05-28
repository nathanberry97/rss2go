# rss2go

![rss2go](./assets/rss2go.png)

> rss2go is a lightweight, self-hosted RSS reader built with Go

## Getting started

If you want to run the application locally ensure you have `sass` and `go`
installed then run the following command:

```
make run
```

Which will allow you to access the website on `localhost:8080`.

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
    <title>rss2go sample opml</title>
    <ownerName>rss2go</ownerName>
  </head>
  <body>
    <outline text="Technology" title="Technology">
      <outline text="AWS" type="rss" xmlUrl="https://aws.amazon.com/blogs/aws/feed/" />
      <outline text="Microsoft" type="rss" xmlUrl="https://devblogs.microsoft.com/landing" />
      <outline text="Cloudflare" type="rss" xmlUrl="https://blog.cloudflare.com/rss" />
      <outline text="Go Blog" type="atom" xmlUrl="http://blog.golang.org/feed.atom" />
      <outline text="Framework" type="rss" xmlUrl="https://frame.work/gb/en/blog.rss" />
      <outline text="BBC" type="rss" xmlUrl="http://feeds.bbci.co.uk/news/technology/rss.xml" />
      <outline text="Technology in government" type="atom" xmlUrl="https://technology.blog.gov.uk/feed/" />
    </outline>
    <outline text="Personal Blogs" title="Personal Blogs">
      <outline text="Nathan Berry" type="rss" xmlUrl="https://nathanberry.co.uk/feed.xml" />
      <outline text="David Heinemeier Hansson" type="atom" xmlUrl="https://world.hey.com/dhh/feed.atom" />
      <outline text="Coding Horror" type="rss" xmlUrl="https://blog.codinghorror.com/rss" />
    </outline>
  </body>
</opml>
```

## API Endpoints

```txt
GET    /
GET    /health-check
GET    /articles/{feedId}
GET    /articles/favourites
GET    /articles/later
GET    /feeds
GET    /partials/articles
GET    /partials/articles/{id}
GET    /partials/feed
POST   /partials/feed
DELETE /partials/feed/{id}
GET    /partials/feed/opml
POST   /partials/feed/opml
GET    /partials/later
POST   /partials/later/{articleId}
DELETE /partials/later/{articleId}
GET    /partials/favourite
POST   /partials/favourite/{id}
DELETE /partials/favourite/{id}
```
