package components

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRenderRSSTemplate(t *testing.T) {
	tempDir := t.TempDir()

	tmplContent := `{{ define "test_template" }}
<opml version="1.0">
  <body>
    {{ range . }}
    <outline text="{{ .Name }}" xmlUrl="{{ .URL }}" type="{{ .Type }}"/>
    {{ end }}
  </body>
</opml>
{{ end }}`

	tmplPath := filepath.Join(tempDir, "test.tmpl")
	err := os.WriteFile(tmplPath, []byte(tmplContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test template: %v", err)
	}

	type testData struct {
		Name string
		URL  string
		Type string
	}

	tests := []struct {
		name     string
		path     string
		tmplName string
		data     any
		wantErr  bool
	}{
		{
			name:     "valid template with data",
			path:     tmplPath,
			tmplName: "test_template",
			data: []testData{
				{Name: "Test Feed", URL: "https://example.com/feed.xml", Type: "rss"},
			},
			wantErr: false,
		},
		{
			name:     "non-existent template",
			path:     "/non/existent/path.tmpl",
			tmplName: "test_template",
			data:     nil,
			wantErr:  true,
		},
		{
			name:     "wrong template name",
			path:     tmplPath,
			tmplName: "nonexistent_template",
			data:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderRSSTemplate(tt.path, tt.tmplName, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderRSSTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) == 0 {
					t.Errorf("RenderRSSTemplate() returned empty byte slice")
				}
				if string(got[:5]) != "<?xml" {
					t.Errorf("RenderRSSTemplate() should start with XML declaration")
				}
			}
		})
	}
}
