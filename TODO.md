# Todo List for rss2go

## Todo

* [ ] Add dockerfile to build and run the application

* [ ] Implement a database sql
  - [ ] Add a database connection to the API framework
  - **Note** going to use PSQL for the database

* [ ] Add REACT frontend to the application
  - [ ] Add a frontend folder to the project
  - [ ] Add a REACT frontend to the frontend folder

* [ ] Favourite articles endpoints
  - [ ] POST favourite articles
  - [ ] GET favourite articles
  - [ ] DELETE favourite articles

## In progress

* [ ] RSS feed endpoints
  - [ ] POST new RSS feeds to the database
  - [ ] GET RSS feeds
  - [ ] GET RSS feed posts from using the RSS feed URL
  - [ ] GET RSS feed posts from all RSS feeds
  - [ ] DELETE RSS feeds from the database

> **NOTE** use YAML file for the time being to store the RSS feeds

## Completed

* [X] Add linting pipeline and precommit
* [X] Setup API framework with go and add health check endpoint
