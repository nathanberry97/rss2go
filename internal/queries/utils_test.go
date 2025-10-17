package queries

import (
	"strings"
	"testing"
)

func TestBuildArticleQuery(t *testing.T) {
	tests := []struct {
		name       string
		conditions []string
		wantSubstr []string
	}{
		{
			name:       "no conditions",
			conditions: []string{},
			wantSubstr: []string{
				"SELECT f.name, a.feed_id, a.id, a.title, a.url, a.published_at",
				"FROM articles a",
				"JOIN feeds f ON a.feed_id = f.id",
				"ORDER BY published_at DESC",
				"LIMIT ? OFFSET ?",
			},
		},
		{
			name:       "single condition",
			conditions: []string{"a.feed_id = ?"},
			wantSubstr: []string{
				"WHERE a.feed_id = ?",
				"ORDER BY published_at DESC",
			},
		},
		{
			name:       "multiple conditions",
			conditions: []string{"a.feed_id = ?", "a.published_at > ?"},
			wantSubstr: []string{
				"WHERE a.feed_id = ? AND a.published_at > ?",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildArticleQuery(tt.conditions...)
			for _, substr := range tt.wantSubstr {
				if !strings.Contains(got, substr) {
					t.Errorf("buildArticleQuery() missing expected substring: %q\nGot: %s", substr, got)
				}
			}
		})
	}
}

func TestBuildArticleTotalQuery(t *testing.T) {
	tests := []struct {
		name       string
		table      string
		conditions []string
		want       string
	}{
		{
			name:       "no conditions",
			table:      "articles",
			conditions: []string{},
			want:       "SELECT COUNT(*) FROM articles",
		},
		{
			name:       "single condition",
			table:      "articles",
			conditions: []string{"feed_id = ?"},
			want:       "SELECT COUNT(*) FROM articles\nWHERE feed_id = ?",
		},
		{
			name:       "multiple conditions",
			table:      "favourites",
			conditions: []string{"feed_id = ?", "created_at > ?"},
			want:       "SELECT COUNT(*) FROM favourites\nWHERE feed_id = ? AND created_at > ?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildArticleTotalQuery(tt.table, tt.conditions...)
			if got != tt.want {
				t.Errorf("buildArticleTotalQuery() = %q, want %q", got, tt.want)
			}
		})
	}
}
