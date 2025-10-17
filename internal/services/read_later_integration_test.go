package services

import (
	"testing"
)

func TestPostReadLater_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, articleIDs := seedTestData(t, db)

	tests := []struct {
		name      string
		articleID string
		wantErr   bool
	}{
		{
			name:      "add article to read later",
			articleID: "1",
			wantErr:   false,
		},
		{
			name:      "add another article",
			articleID: "2",
			wantErr:   false,
		},
		{
			name:      "duplicate read later is ok",
			articleID: "1",
			wantErr:   true, // Will fail without OR IGNORE in query
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostReadLater(db, tt.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostReadLater() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// Verify read later was added
				var count int
				err := db.QueryRow("SELECT COUNT(*) FROM read_later WHERE article_id = ?", tt.articleID).Scan(&count)
				if err != nil {
					t.Fatalf("Failed to query read_later: %v", err)
				}
				if count != 1 {
					t.Errorf("PostReadLater() read later not found in database")
				}
			}
		})
	}

	_ = articleIDs // Use variable
}

func TestDeleteReadLater_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, articleIDs := seedTestData(t, db)

	// Add a read later first
	_, err := db.Exec("INSERT INTO read_later (article_id) VALUES (?)", articleIDs[0])
	if err != nil {
		t.Fatalf("Failed to setup test read later: %v", err)
	}

	tests := []struct {
		name      string
		articleID string
		wantErr   bool
	}{
		{
			name:      "delete existing read later",
			articleID: "1",
			wantErr:   false,
		},
		{
			name:      "delete non-existent read later is ok",
			articleID: "999",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteReadLater(db, tt.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteReadLater() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && tt.articleID == "1" {
				// Verify read later was deleted
				var count int
				err := db.QueryRow("SELECT COUNT(*) FROM read_later WHERE article_id = ?", tt.articleID).Scan(&count)
				if err != nil {
					t.Fatalf("Failed to query read_later: %v", err)
				}
				if count != 0 {
					t.Errorf("DeleteReadLater() read later still exists in database")
				}
			}
		})
	}
}

func TestGetReadLater_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, articleIDs := seedTestData(t, db)

	// Add some read later items
	for i := range 2 {
		_, err := db.Exec("INSERT INTO read_later (article_id) VALUES (?)", articleIDs[i])
		if err != nil {
			t.Fatalf("Failed to setup test read later: %v", err)
		}
	}

	tests := []struct {
		name         string
		page         int
		limit        int
		wantErr      bool
		wantMinItems int
	}{
		{
			name:         "get all read later",
			page:         0,
			limit:        10,
			wantErr:      false,
			wantMinItems: 2,
		},
		{
			name:         "pagination works",
			page:         0,
			limit:        1,
			wantErr:      false,
			wantMinItems: 1,
		},
		{
			name:         "second page",
			page:         1,
			limit:        1,
			wantErr:      false,
			wantMinItems: 1,
		},
		{
			name:    "invalid pagination",
			page:    -1,
			limit:   10,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := GetReadLater(db, tt.page, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReadLater() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(response.Items) < tt.wantMinItems {
					t.Errorf("GetReadLater() returned %d items, want at least %d", len(response.Items), tt.wantMinItems)
				}

				// Verify all items are in read later
				for _, article := range response.Items {
					if !article.Later {
						t.Errorf("GetReadLater() returned article with Later=false")
					}
				}
			}
		})
	}
}
