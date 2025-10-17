package css

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHashCSSFile(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name        string
		setup       func() (string, string)
		wantErr     bool
		wantPattern string
	}{
		{
			name: "valid CSS file",
			setup: func() (string, string) {
				tempFile := "temp.css"
				content := []byte("body { color: red; }")
				err := os.WriteFile(filepath.Join(tempDir, tempFile), content, 0644)
				if err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
				return tempDir, tempFile
			},
			wantErr:     false,
			wantPattern: "style-",
		},
		{
			name: "non-existent file",
			setup: func() (string, string) {
				return tempDir, "nonexistent.css"
			},
			wantErr:     true,
			wantPattern: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputDir, tempFile := tt.setup()
			got, err := HashCSSFile(outputDir, tempFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashCSSFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) == 0 {
					t.Errorf("HashCSSFile() returned empty filename")
				}
				if tt.wantPattern != "" && len(got) > 0 {
					if got[:6] != tt.wantPattern {
						t.Errorf("HashCSSFile() = %v, want pattern %v*", got, tt.wantPattern)
					}
				}
			}
		})
	}
}
