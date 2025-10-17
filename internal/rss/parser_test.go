package rss

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	tests := []struct {
		name         string
		dateStr      string
		inputFormat  string
		outputFormat string
		want         string
		wantErr      bool
	}{
		{
			name:         "RFC1123 to DateTime",
			dateStr:      "Mon, 02 Jan 2006 15:04:05 MST",
			inputFormat:  time.RFC1123,
			outputFormat: time.DateTime,
			want:         "2006-01-02 15:04:05",
			wantErr:      false,
		},
		{
			name:         "RFC3339 to DateTime",
			dateStr:      "2006-01-02T15:04:05Z",
			inputFormat:  time.RFC3339,
			outputFormat: time.DateTime,
			want:         "2006-01-02 15:04:05",
			wantErr:      false,
		},
		{
			name:         "invalid date",
			dateStr:      "invalid",
			inputFormat:  time.RFC3339,
			outputFormat: time.DateTime,
			want:         "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatDate(tt.dateStr, tt.inputFormat, tt.outputFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("formatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsePubDate(t *testing.T) {
	tests := []struct {
		name    string
		pubDate string
		wantErr bool
	}{
		{
			name:    "RFC1123 format",
			pubDate: "Mon, 02 Jan 2006 15:04:05 MST",
			wantErr: false,
		},
		{
			name:    "RFC1123Z format",
			pubDate: "Mon, 02 Jan 2006 15:04:05 -0700",
			wantErr: false,
		},
		{
			name:    "RFC3339 format",
			pubDate: "2006-01-02T15:04:05Z",
			wantErr: false,
		},
		{
			name:    "DateTime format",
			pubDate: "2006-01-02 15:04:05",
			wantErr: false,
		},
		{
			name:    "invalid format",
			pubDate: "not a valid date",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsePubDate(tt.pubDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePubDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("parsePubDate() returned empty string for valid input")
			}
		})
	}
}
