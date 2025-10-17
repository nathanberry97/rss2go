package services

import (
	"testing"

	"github.com/nathanberry97/rss2go/internal/schema"
)

func TestPostFeedOpml_ParseValidation(t *testing.T) {
	tests := []struct {
		name     string
		opmlData []byte
		wantErr  bool
	}{
		{
			name:     "invalid XML",
			opmlData: []byte(`not valid xml`),
			wantErr:  true,
		},
		{
			name: "empty feeds",
			opmlData: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<opml version="1.0">
  <body>
  </body>
</opml>`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostFeedOpml(nil, tt.opmlData)
			hasErr := err != nil
			if hasErr != tt.wantErr {
				t.Errorf("PostFeedOpml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestImportSingleFeed(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tests := []struct {
		name    string
		feedURL string
		wantErr bool
	}{
		{
			name:    "empty URL",
			feedURL: "",
			wantErr: true,
		},
		{
			name:    "whitespace URL",
			feedURL: "   ",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := importSingleFeed(db, tt.feedURL)
			hasErr := err != nil
			if hasErr != tt.wantErr {
				t.Errorf("importSingleFeed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostFeed_Validation(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tests := []struct {
		name     string
		postBody schema.RssPostBody
		wantErr  bool
	}{
		{
			name: "empty URL",
			postBody: schema.RssPostBody{
				URL: "",
			},
			wantErr: true,
		},
		{
			name: "whitespace URL",
			postBody: schema.RssPostBody{
				URL: "   ",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostFeed(db, tt.postBody)
			hasErr := err != nil
			if hasErr != tt.wantErr {
				t.Errorf("PostFeed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
