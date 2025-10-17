package services

import (
	"testing"
)

func TestGetReadLater_ValidationOnly(t *testing.T) {
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
			_, err := GetReadLater(nil, tt.page, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReadLater() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
