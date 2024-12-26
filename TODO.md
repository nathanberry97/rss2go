# Todo List for rss2go

## Todo

* [ ] Add dockerfile to build and run the application

* [ ] Add REACT frontend to the application
  - [ ] Add a frontend folder to the project
  - [ ] Add a REACT frontend to the frontend folder

* [ ] Favourite articles endpoints
  - [ ] POST favourite articles
  - [ ] GET favourite articles
  - [ ] DELETE favourite articles

* [ ] RSS feed endpoints
  - [ ] POST new RSS feeds to the database
  - [ ] GET RSS feeds
  - [ ] GET RSS feed posts from using the RSS feed URL
  - [ ] GET RSS feed posts from all RSS feeds
  - [ ] DELETE RSS feeds from the database

## In progress

* [ ] Implement a database sql
  - [ ] Add Liquibase to the project to manage database
  - [ ] Add postgres to the dockerfile
  - [ ] Create a database schema for the project
  - [ ] Add a database connection to the API framework

> **Note** going to use PSQL for the database

## Completed

* [X] Add linting pipeline and precommit
* [X] Setup API framework with go and add health check endpoint
