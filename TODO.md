# Todo List for rss2go

## Todo

* [ ] Documentation site
  - [ ] Create site
  - [ ] Host site in AWS or Github pages?

* [ ] GET articles by feed
  - [ ] GET
  - [ ] Frontend

* [ ] Add loading animations

* [ ] Refine scss and responsive design

* [ ] Read later
  - [ ] POST
  - [ ] GET
  - [ ] DELETE
  - [ ] Frontend
  - UI icon bookmark

* [ ] Favourite articles
  - [ ] POST
  - [ ] GET
  - [ ] DELETE
  - [ ] Frontend
  - UI icon heart

* [ ] Add better error handling
    - [ ] Backend to return HTML to the frontend
    - [ ] Frontend to show the error in the UI

* [ ] Upload to docker hub and release pipeline
    - [ ] Works on v0.0.0 branches
    - [ ] Pushes the build to docker hub
    - [ ] Creates a release within the repo itself

* [ ] Add dockerfile to build and run the application

* [ ] Add support for OPML
    - [ ] POST
    - [ ] GET
    - [ ] Frontend

## In progress

* [ ] Implement unit tests for the application
    - [ ] worker
    - [ ] rss
    - [ ] services
    - [ ] utils

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
* [X] Improve the rss reader function
  - [X] Improve the general function
  - [X] Make it work with .atom feeds too
* [X] Integrate go worker pool to update feeds in the background
    - [X] Run on a schedule (Every 1 hours perhaps?)
    - [X] Run when server starts up
    - [X] Updates the articles table
* [X] Improve scss and js into the project
  - [X] Refine current scss and js
  - [X] Make responsive (ensure it looks fine on monitor and laptop screen)
* [X] Migrate over from `colly` to `gofeed` for rss parser
