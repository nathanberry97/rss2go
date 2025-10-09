package main

import (
	"log"

	"github.com/nathanberry97/rss2go/internal/css"
	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/jobs"
	"github.com/nathanberry97/rss2go/internal/routes"
)

func main() {
	// Set environment variables
	port := "8080"

	// Set up the database
	db := database.DatabaseConnection()
	defer db.Close()

	err := database.InitializeDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	// Set up worker pool to refresh the feeds async
	go jobs.SyncFeeds(5)

	// Hashed CSS file name
	hashedCSS, err := css.HashCSSFile("web/static/css", "style.tmp.css")
	if err != nil {
		log.Fatalf("CSS hashing failed: %v", err)
	}

	// Start the server
	router := routes.InitialiseRouter(hashedCSS)
	router.Run(":" + port)
}
