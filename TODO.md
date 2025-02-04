# Todo List for rss2go

## Todo

* [ ] Upload to docker hub and release pipeline

* [ ] Documentation site

* [ ] Add dockerfile to build and run the application

* [ ] Improve scss and js into the project
  - [ ] Refine current scss and js
  - [ ] Add loading animation
  - [ ] Make responsive

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

* [ ] Implement unit tests for the application

* [ ] Add better error handling

## In progress

* [ ] Improve the rss reader function
  - [ ] Improve the general function
  - [ ] Make it work on a schedule to allow new posts to be added
  - [ ] Make it work with .atom feeds too

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
* [X] Add scss and js into the project
