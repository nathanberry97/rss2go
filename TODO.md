# Todo List for rss2go

## Todo

* [ ] Add dockerfile to build and run the application

* [ ] Add HTMX to the frontend
  - [ ] Add a frontend folder to the project
  - [ ] Implement HTMX to the frontend to consume the API

* [ ] Favourite articles endpoints
  - [ ] POST favourite articles
  - [ ] GET favourite articles
  - [ ] DELETE favourite articles

## In progress

* [ ] RSS Article endpoints
  - [ ] Update POST new RSS feed to add articles to the database for the feed
  - [ ] GET RSS articles
  - [ ] DELETE RSS articles

## Completed

* [X] Add linting pipeline and precommit
* [X] Setup API framework with go and add health check endpoint
* [X] Implement a database sql
* [X] Add a database connection to the application
* [X] RSS feed endpoints
  - [X] POST new RSS feeds to the database
  - [X] GET RSS feeds
  - [X] DELETE RSS feeds from the database
