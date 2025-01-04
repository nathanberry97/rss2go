--liquibase formatted sql

--changeset nathan:1-create-feeds-table
CREATE TABLE feeds (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL
);
--rollback DROP TABLE feeds;

--changeset nathan:2-create-articles-table
CREATE TABLE articles (
  id SERIAL PRIMARY KEY,
  feed_id INT NOT NULL,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  description TEXT,
  FOREIGN KEY (feed_id) REFERENCES feeds(id)
);
--rollback DROP TABLE articles;

--changeset nathan:3-create-favorites-table
CREATE TABLE favorites (
  id SERIAL PRIMARY KEY,
  article_id INT NOT NULL,
  FOREIGN KEY (article_id) REFERENCES articles(id)
);
--rollback DROP TABLE favorites;
