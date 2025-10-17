package services

import (
	"testing"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func TestGetArticles_ValidationOnly(t *testing.T) {
	tests := []struct {
		name    string
		page    int
		limit   int
		wantErr bool
	}{
		{
			name:    "negative page",
			page:    -1,
			limit:   10,
			wantErr: true,
		},
		{
			name:    "zero limit",
			page:    0,
			limit:   0,
			wantErr: true,
		},
		{
			name:    "negative limit",
			page:    0,
			limit:   -5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetArticles(nil, tt.page, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetArticlesByFeedId_ValidationOnly(t *testing.T) {
	tests := []struct {
		name    string
		page    int
		limit   int
		id      int
		wantErr bool
	}{
		{
			name:    "negative page",
			page:    -1,
			limit:   10,
			id:      1,
			wantErr: true,
		},
		{
			name:    "zero limit",
			page:    0,
			limit:   0,
			id:      1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetArticlesByFeedId(nil, tt.page, tt.limit, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticlesByFeedId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertArticles(t *testing.T) {
	tests := []struct {
		name     string
		articles []schema.RssItem
		feedID   int64
		wantErr  bool
	}{
		{
			name:     "empty articles",
			articles: []schema.RssItem{},
			feedID:   1,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InsertArticles(nil, tt.articles, tt.feedID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertArticles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
