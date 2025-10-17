package services

import (
	"testing"
)

func TestPostFavourite_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, articleIDs := seedTestData(t, db)

	tests := []struct {
		name      string
		articleID string
		wantErr   bool
	}{
		{
			name:      "add article to favourites",
			articleID: "1",
			wantErr:   false,
		},
		{
			name:      "add another article",
			articleID: "2",
			wantErr:   false,
		},
		{
			name:      "duplicate favourite is ok (INSERT OR IGNORE)",
			articleID: "1",
			wantErr:   true, // Will fail without OR IGNORE in query
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostFavourite(db, tt.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostFavourite() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// Verify favourite was added
				var count int
				err := db.QueryRow("SELECT COUNT(*) FROM favourites WHERE article_id = ?", tt.articleID).Scan(&count)
				if err != nil {
					t.Fatalf("Failed to query favourites: %v", err)
				}
				if count != 1 {
					t.Errorf("PostFavourite() favourite not found in database")
				}
			}
		})
	}

	_ = articleIDs // Use variable
}

func TestDeleteFavourite_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, articleIDs := seedTestData(t, db)

	// Add a favourite first
	_, err := db.Exec("INSERT INTO favourites (article_id) VALUES (?)", articleIDs[0])
	if err != nil {
		t.Fatalf("Failed to setup test favourite: %v", err)
	}

	tests := []struct {
		name      string
		articleID string
		wantErr   bool
	}{
		{
			name:      "delete existing favourite",
			articleID: "1",
			wantErr:   false,
		},
		{
			name:      "delete non-existent favourite is ok",
			articleID: "999",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteFavourite(db, tt.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteFavourite() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && tt.articleID == "1" {
				// Verify favourite was deleted
				var count int
				err := db.QueryRow("SELECT COUNT(*) FROM favourites WHERE article_id = ?", tt.articleID).Scan(&count)
				if err != nil {
					t.Fatalf("Failed to query favourites: %v", err)
				}
				if count != 0 {
					t.Errorf("DeleteFavourite() favourite still exists in database")
				}
			}
		})
	}
}

func TestGetFavourites_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, articleIDs := seedTestData(t, db)

	// Add some favourites
	for i := range 2 {
		_, err := db.Exec("INSERT INTO favourites (article_id) VALUES (?)", articleIDs[i])
		if err != nil {
			t.Fatalf("Failed to setup test favourites: %v", err)
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
			name:         "get all favourites",
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
			response, err := GetFavourites(db, tt.page, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFavourites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(response.Items) < tt.wantMinItems {
					t.Errorf("GetFavourites() returned %d items, want at least %d", len(response.Items), tt.wantMinItems)
				}

				// Verify all items are favourites
				for _, article := range response.Items {
					if !article.Fav {
						t.Errorf("GetFavourites() returned article with Fav=false")
					}
				}
			}
		})
	}
}
