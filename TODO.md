# Todo List for rss2go

## Todo

* [ ] Documentation site (probably be its own git repo)
  - [ ] Create site
  - [ ] Host site in AWS or Github pages?

* [ ] Upload to docker hub and release pipeline
    - [ ] Works on v1.0.0 branches
    - [ ] Pushes the build to docker hub
    - [ ] Creates a release within the repo itself

* [ ] Add better error handling
    - [ ] Backend to return HTML to the frontend
    - [ ] Frontend to show the error in the UI

* [ ] Add dockerfile to build and run the application

* [ ] Implement unit tests for the application
    - [ ] worker
    - [ ] rss
    - [ ] services
    - [ ] utils

* [ ] Add support for OPML (Inside of settings page)
    - [ ] POST
    - [ ] GET
    - [ ] Frontend

## In progress

* [ ] Refine scss and responsive design
    - [X] Add scss compile into the build process and add hash
    - [ ] Redo article layout
        - [ ] Add Read Later button
        - [ ] Add Favourite button
    - [ ] Add loading animations
    - [ ] Add light theme and dark theme

## Completed

* [X] Add linting pipeline and precommit
* [X] Setup API framework with go and add health check endpoint
* [X] Implement a database sql
* [X] Add a database connection to the application
* [X] RSS feed endpoints
* [X] RSS Article endpoints
* [X] Implement seeding script
* [X] Add HTMX to the frontend
* [X] Add scss and js into the project
* [X] Improve the rss reader function
* [X] Integrate go worker pool to update feeds in the background
* [X] Improve scss and js into the project
* [X] Migrate over from `colly` to `gofeed` for rss parser
* [X] GET articles by feed
* [X] Create placeholder pages
* [X] Clean up templates
* [X] Read later
* [X] Favourite articles
