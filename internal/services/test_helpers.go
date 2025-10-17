package services

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

// setupTestDB creates an in-memory SQLite database for testing
func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Enable foreign keys
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		t.Fatalf("Failed to enable foreign keys: %v", err)
	}

	// Create tables
	schema := `
	CREATE TABLE IF NOT EXISTS feeds (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  url TEXT NOT NULL UNIQUE
	);

	CREATE TABLE IF NOT EXISTS articles (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  feed_id INTEGER NOT NULL,
	  title TEXT NOT NULL,
	  url TEXT NOT NULL UNIQUE,
	  published_at DATETIME,
	  FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS favourites (
	  article_id INTEGER PRIMARY KEY,
	  FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS read_later (
	  article_id INTEGER PRIMARY KEY,
	  FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	return db
}

// seedTestData inserts test data into the database
func seedTestData(t *testing.T, db *sql.DB) (feedID int64, articleIDs []int64) {
	t.Helper()

	// Insert a test feed
	result, err := db.Exec("INSERT INTO feeds (name, url) VALUES (?, ?)", "Test Feed", "https://example.com/feed.xml")
	if err != nil {
		t.Fatalf("Failed to insert test feed: %v", err)
	}

	feedID, err = result.LastInsertId()
	if err != nil {
		t.Fatalf("Failed to get feed ID: %v", err)
	}

	// Insert test articles with recent dates (within last 31 days for GetArticles)
	articles := []struct {
		title       string
		url         string
		publishedAt string
	}{
		{"Article 1", "https://example.com/article1", "datetime('now', '-1 day')"},
		{"Article 2", "https://example.com/article2", "datetime('now', '-2 days')"},
		{"Article 3", "https://example.com/article3", "datetime('now', '-3 days')"},
	}

	for _, article := range articles {
		result, err := db.Exec(
			"INSERT INTO articles (feed_id, title, url, published_at) VALUES (?, ?, ?, "+article.publishedAt+")",
			feedID, article.title, article.url,
		)
		if err != nil {
			t.Fatalf("Failed to insert test article: %v", err)
		}

		articleID, err := result.LastInsertId()
		if err != nil {
			t.Fatalf("Failed to get article ID: %v", err)
		}
		articleIDs = append(articleIDs, articleID)
	}

	return feedID, articleIDs
}
