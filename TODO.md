# Todo List for rss2go

## Todo

* [ ] Add dockerfile to build and run the application

* [ ] GET articles by feed
  - [ ] GET
  - [ ] Frontend

* [ ] Read later
  - [ ] POST
  - [ ] GET
  - [ ] DELETE
  - [ ] Frontend

* [ ] Favourite articles
  - [ ] POST
  - [ ] GET
  - [ ] DELETE
  - [ ] Frontend

* [ ] Add better error handling

* [ ] Improve the rss reader function
  - [ ] Make it work on a schedule to allow new posts to be added
  - [ ] Make it work with .atom feeds too

* [ ] Implement unit tests for the application

## In progress

* [ ] Add scss and js into the project
  - Add css to improve ui [ ]
  - Add alpine or implement code to clear form after sending the request [ ]
  - Implement loading bar on feeds page [ ]

## Completed

* [X] Add linting pipeline and precommit
* [X] Setup API framework with go and add health check endpoint
* [X] Implement a database sql
* [X] Add a database connection to the application
* [X] RSS feed endpoints
  - [X] POST
  - [X] GET
  - [X] DELETE
* [X] RSS Article endpoints
  - [X] GET
  - [X] POST (Update RSS feed post to do this)
* [X] Implement seeding script
* [X] Add HTMX to the frontend
