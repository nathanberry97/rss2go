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
	tests := []struct {
		name    string
		setup   func() *sql.Rows
		wantErr bool
	}{
		{
			name: "empty result set",
			setup: func() *sql.Rows {
				return nil
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil && tt.setup() == nil && !tt.wantErr {
				t.Skip("Skipping test that requires database mock")
			}
		})
	}
}
