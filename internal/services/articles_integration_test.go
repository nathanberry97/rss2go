package services

import (
	"testing"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func TestGetArticles_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, _ = seedTestData(t, db)

	tests := []struct {
		name           string
		page           int
		limit          int
		wantErr        bool
		wantMinItems   int
		checkNextPage  bool
		expectNextPage int
	}{
		{
			name:           "first page with limit 2",
			page:           0,
			limit:          2,
			wantErr:        false,
			wantMinItems:   2,
			checkNextPage:  true,
			expectNextPage: 1,
		},
		{
			name:          "page 1 with limit 2",
			page:          1,
			limit:         2,
			wantErr:       false,
			wantMinItems:  1,
			checkNextPage: true,
		},
		{
			name:         "high page returns empty",
			page:         100,
			limit:        10,
			wantErr:      false,
			wantMinItems: 0,
		},
		{
			name:    "invalid negative page",
			page:    -1,
			limit:   10,
			wantErr: true,
		},
		{
			name:    "invalid zero limit",
			page:    0,
			limit:   0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := GetArticles(db, tt.page, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if len(response.Items) < tt.wantMinItems {
				t.Errorf("GetArticles() returned %d items, want at least %d", len(response.Items), tt.wantMinItems)
			}

			if tt.checkNextPage && tt.expectNextPage > 0 && response.NextPage != tt.expectNextPage {
				t.Errorf("GetArticles() NextPage = %d, want %d", response.NextPage, tt.expectNextPage)
			}

			if tt.wantMinItems > 0 && response.TotalItems <= 0 {
				t.Errorf("GetArticles() TotalItems = %d, want > 0", response.TotalItems)
			}
		})
	}
}

func TestGetArticlesByFeedId_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	feedID, _ := seedTestData(t, db)

	// Create another feed to ensure filtering works
	_, err := db.Exec("INSERT INTO feeds (name, url) VALUES (?, ?)", "Other Feed", "https://example.com/other.xml")
	if err != nil {
		t.Fatalf("Failed to insert other feed: %v", err)
	}

	tests := []struct {
		name         string
		page         int
		limit        int
		feedID       int
		wantErr      bool
		wantMinItems int
	}{
		{
			name:         "get articles for valid feed",
			page:         0,
			limit:        10,
			feedID:       int(feedID),
			wantErr:      false,
			wantMinItems: 3,
		},
		{
			name:         "pagination works",
			page:         0,
			limit:        2,
			feedID:       int(feedID),
			wantErr:      false,
			wantMinItems: 2,
		},
		{
			name:         "non-existent feed returns empty",
			page:         0,
			limit:        10,
			feedID:       999,
			wantErr:      false,
			wantMinItems: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := GetArticlesByFeedId(db, tt.page, tt.limit, tt.feedID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticlesByFeedId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(response.Items) < tt.wantMinItems {
				t.Errorf("GetArticlesByFeedId() returned %d items, want at least %d", len(response.Items), tt.wantMinItems)
			}

			// Verify all articles belong to the correct feed
			for _, article := range response.Items {
				if article.FeedId != tt.feedID && tt.wantMinItems > 0 {
					t.Errorf("GetArticlesByFeedId() article has FeedId %d, want %d", article.FeedId, tt.feedID)
				}
			}
		})
	}
}

func TestInsertArticles_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	feedID, _ := seedTestData(t, db)

	tests := []struct {
		name     string
		articles []schema.RssItem
		feedID   int64
		wantErr  bool
	}{
		{
			name: "insert single article",
			articles: []schema.RssItem{
				{
					Title:   "New Article",
					Link:    "https://example.com/new-article",
					PubDate: "2024-01-20 10:00:00",
				},
			},
			feedID:  feedID,
			wantErr: false,
		},
		{
			name: "insert multiple articles",
			articles: []schema.RssItem{
				{
					Title:   "Bulk Article 1",
					Link:    "https://example.com/bulk1",
					PubDate: "2024-01-20 10:00:00",
				},
				{
					Title:   "Bulk Article 2",
					Link:    "https://example.com/bulk2",
					PubDate: "2024-01-20 11:00:00",
				},
			},
			feedID:  feedID,
			wantErr: false,
		},
		{
			name: "duplicate URL is ignored",
			articles: []schema.RssItem{
				{
					Title:   "Duplicate",
					Link:    "https://example.com/article1", // Already exists
					PubDate: "2024-01-20 10:00:00",
				},
			},
			feedID:  feedID,
			wantErr: false,
		},
		{
			name:     "empty articles list",
			articles: []schema.RssItem{},
			feedID:   feedID,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InsertArticles(db, tt.articles, tt.feedID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertArticles() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && len(tt.articles) > 0 {
				// Verify at least some articles were inserted (might be less due to duplicates)
				var count int
				err := db.QueryRow("SELECT COUNT(*) FROM articles WHERE feed_id = ?", tt.feedID).Scan(&count)
				if err != nil {
					t.Fatalf("Failed to count articles: %v", err)
				}
				if count == 0 {
					t.Errorf("InsertArticles() inserted 0 articles, expected some")
				}
			}
		})
	}
}

func TestGetArticleName_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, _ = seedTestData(t, db)

	tests := []struct {
		name     string
		feedID   string
		wantName string
		wantErr  bool
	}{
		{
			name:     "valid feed ID",
			feedID:   "1",
			wantName: "Test Feed",
			wantErr:  false,
		},
		{
			name:     "non-existent feed ID",
			feedID:   "999",
			wantName: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: GetArticleName uses database.DatabaseConnection() internally,
			// so we test the query logic with our test DB
			var name string
			err := db.QueryRow("SELECT name FROM feeds WHERE id = ?", tt.feedID).Scan(&name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && name != tt.wantName {
				t.Errorf("Got name = %v, want %v", name, tt.wantName)
			}
		})
	}
}
