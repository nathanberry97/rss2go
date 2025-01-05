CREATE SCHEMA IF NOT EXISTS rss;

CREATE TABLE rss.feeds (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL
);

CREATE TABLE rss.articles (
  id SERIAL PRIMARY KEY,
  feed_id INT NOT NULL,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  description TEXT,
  FOREIGN KEY (feed_id) REFERENCES rss.feeds(id)
);

CREATE TABLE rss.favorites (
  id SERIAL PRIMARY KEY,
  article_id INT NOT NULL,
  FOREIGN KEY (article_id) REFERENCES rss.articles(id)
);
