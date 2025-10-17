package queries

import (
	"strings"
	"testing"
)

func TestInsertArticlesQuery(t *testing.T) {
	tests := []struct {
		name          string
		num           int
		wantPrefix    string
		wantNumValues int
	}{
		{
			name:          "single article",
			num:           1,
			wantPrefix:    "INSERT OR IGNORE INTO articles (feed_id, title, url, published_at) VALUES",
			wantNumValues: 1,
		},
		{
			name:          "multiple articles",
			num:           3,
			wantPrefix:    "INSERT OR IGNORE INTO articles (feed_id, title, url, published_at) VALUES",
			wantNumValues: 3,
		},
		{
			name:          "zero articles",
			num:           0,
			wantPrefix:    "INSERT OR IGNORE INTO articles (feed_id, title, url, published_at) VALUES",
			wantNumValues: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertArticlesQuery(tt.num)
			if !strings.HasPrefix(got, tt.wantPrefix) {
				t.Errorf("InsertArticlesQuery() prefix = %q, want prefix %q", got, tt.wantPrefix)
			}

			placeholderCount := strings.Count(got, "(?, ?, ?, ?)")
			if placeholderCount != tt.wantNumValues {
				t.Errorf("InsertArticlesQuery() has %d value sets, want %d", placeholderCount, tt.wantNumValues)
			}
		})
	}
}
