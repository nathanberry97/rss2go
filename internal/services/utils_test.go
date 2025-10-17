package services

import (
	"database/sql"
	"testing"
	"time"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func TestInferFeedType(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "atom with .atom extension",
			url:  "https://example.com/feed.atom",
			want: "atom",
		},
		{
			name: "atom with /atom in path",
			url:  "https://example.com/atom",
			want: "atom",
		},
		{
			name: "atom with format=atom parameter",
			url:  "https://example.com/feed?format=atom",
			want: "atom",
		},
		{
			name: "rss default",
			url:  "https://example.com/feed.xml",
			want: "rss",
		},
		{
			name: "rss with /rss in path",
			url:  "https://example.com/rss",
			want: "rss",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferFeedType(tt.url)
			if got != tt.want {
				t.Errorf("inferFeedType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractFeedUrls(t *testing.T) {
	tests := []struct {
		name     string
		outlines []schema.OpmlOutline
		want     []string
	}{
		{
			name: "single outline with URL",
			outlines: []schema.OpmlOutline{
				{XMLURL: "https://example.com/feed.xml"},
			},
			want: []string{"https://example.com/feed.xml"},
		},
		{
			name: "nested outlines",
			outlines: []schema.OpmlOutline{
				{
					XMLURL: "https://example1.com/feed.xml",
					Outlines: []schema.OpmlOutline{
						{XMLURL: "https://example2.com/feed.xml"},
					},
				},
			},
			want: []string{"https://example1.com/feed.xml", "https://example2.com/feed.xml"},
		},
		{
			name: "outline without URL",
			outlines: []schema.OpmlOutline{
				{Text: "Category", XMLURL: ""},
			},
			want: []string{},
		},
		{
			name:     "empty outlines",
			outlines: []schema.OpmlOutline{},
			want:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractFeedUrls(tt.outlines)
			if len(got) != len(tt.want) {
				t.Errorf("extractFeedUrls() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("extractFeedUrls()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFormatTime(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name        string
		dateStr     string
		inputFormat string
		want        string
		wantErr     bool
	}{
		{
			name:        "1 hour ago",
			dateStr:     now.Add(-1 * time.Hour).Format(time.RFC3339),
			inputFormat: time.RFC3339,
			want:        "1 hour ago",
			wantErr:     false,
		},
		{
			name:        "5 hours ago",
			dateStr:     now.Add(-5 * time.Hour).Format(time.RFC3339),
			inputFormat: time.RFC3339,
			want:        "5 hours ago",
			wantErr:     false,
		},
		{
			name:        "1 day ago",
			dateStr:     now.Add(-24 * time.Hour).Format(time.RFC3339),
			inputFormat: time.RFC3339,
			want:        "1 day ago",
			wantErr:     false,
		},
		{
			name:        "5 days ago",
			dateStr:     now.Add(-5 * 24 * time.Hour).Format(time.RFC3339),
			inputFormat: time.RFC3339,
			want:        "5 days ago",
			wantErr:     false,
		},
		{
			name:        "1 month ago",
			dateStr:     now.Add(-30 * 24 * time.Hour).Format(time.RFC3339),
			inputFormat: time.RFC3339,
			want:        "1 month ago",
			wantErr:     false,
		},
		{
			name:        "3 months ago",
			dateStr:     now.Add(-90 * 24 * time.Hour).Format(time.RFC3339),
			inputFormat: time.RFC3339,
			want:        "3 months ago",
			wantErr:     false,
		},
		{
			name:        "1 year ago",
			dateStr:     now.Add(-365 * 24 * time.Hour).Format(time.RFC3339),
			inputFormat: time.RFC3339,
			want:        "1 year ago",
			wantErr:     false,
		},
		{
			name:        "2 years ago",
			dateStr:     now.Add(-730 * 24 * time.Hour).Format(time.RFC3339),
			inputFormat: time.RFC3339,
			want:        "2 years ago",
			wantErr:     false,
		},
		{
			name:        "invalid date",
			dateStr:     "invalid",
			inputFormat: time.RFC3339,
			want:        "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatTime(tt.dateStr, tt.inputFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("formatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseOpml(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{
			name: "valid OPML",
			data: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<opml version="1.0">
  <body>
    <outline text="Feed 1" xmlUrl="https://example.com/feed.xml"/>
  </body>
</opml>`),
			wantErr: false,
		},
		{
			name:    "invalid XML",
			data:    []byte(`not valid xml`),
			wantErr: true,
		},
		{
			name:    "empty data",
			data:    []byte{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOpml(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOpml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("parseOpml() returned nil for valid input")
			}
		})
	}
}

func TestFormatArticles(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tests := []struct {
		name         string
		setup        func() *sql.Rows
		wantErr      bool
		wantArticles int
	}{
		{
			name: "empty result set",
			setup: func() *sql.Rows {
				// Return an empty result set
				rows, err := db.Query("SELECT f.name, a.feed_id, a.id, a.title, a.url, a.published_at, 0 AS is_fav, 0 AS is_read_later FROM articles a JOIN feeds f ON a.feed_id = f.id WHERE 1=0")
				if err != nil {
					t.Fatalf("Failed to create empty result set: %v", err)
				}
				return rows
			},
			wantErr:      false,
			wantArticles: 0,
		},
		{
			name: "valid result set with articles",
			setup: func() *sql.Rows {
				// Insert test data
				_, err := db.Exec("INSERT INTO feeds (name, url) VALUES (?, ?)", "Test Feed", "https://example.com/feed.xml")
				if err != nil {
					t.Fatalf("Failed to insert test feed: %v", err)
				}

				_, err = db.Exec("INSERT INTO articles (feed_id, title, url, published_at) VALUES (?, ?, ?, datetime('now', '-1 hour'))", 1, "Test Article", "https://example.com/article1")
				if err != nil {
					t.Fatalf("Failed to insert test article: %v", err)
				}

				rows, err := db.Query(`
					SELECT f.name, a.feed_id, a.id, a.title, a.url, a.published_at,
						0 AS is_fav, 0 AS is_read_later
					FROM articles a
					JOIN feeds f ON a.feed_id = f.id
				`)
				if err != nil {
					t.Fatalf("Failed to query articles: %v", err)
				}
				return rows
			},
			wantErr:      false,
			wantArticles: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := tt.setup()
			if rows != nil {
				defer rows.Close()
			}

			articles, err := formatArticles(rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(articles) != tt.wantArticles {
				t.Errorf("formatArticles() returned %d articles, want %d", len(articles), tt.wantArticles)
			}

			// For non-empty results, verify the article structure
			if tt.wantArticles > 0 {
				for _, article := range articles {
					if article.FeedName == "" {
						t.Errorf("formatArticles() article has empty FeedName")
					}
					if article.Title == "" {
						t.Errorf("formatArticles() article has empty Title")
					}
					if article.PubDate == "" {
						t.Errorf("formatArticles() article has empty PubDate")
					}
					// Check that the time is formatted as relative time
					if article.PubDate != "1 hour ago" {
						t.Logf("formatArticles() PubDate = %s (expected relative time format)", article.PubDate)
					}
				}
			}
		})
	}
}
