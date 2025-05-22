package main

import (
	"log"
	"os"

	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/routes"
	"github.com/nathanberry97/rss2go/internal/scss"
	"github.com/nathanberry97/rss2go/internal/utils"
	"github.com/nathanberry97/rss2go/internal/worker"
)

func main() {
	// Set environment variables
	utils.SetEnv(".env")
	port := os.Getenv("PORT")
	hostName := os.Getenv("HOST_NAME")

	// Set up the database
	db := database.DatabaseConnection()
	defer db.Close()

	err := database.InitializeDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	// Set up worker pool to refresh the feeds async
	go worker.ScheduleFeedUpdates(5)

	// Compile SCSS
	cssFile, err := scss.CompileSCSS("web/scss/main.scss", "web/static/css/")
	if err != nil {
		log.Fatal(err)
	}

	// Start the server
	router := routes.InitialiseRouter(cssFile)
	router.Run(hostName + ":" + port)
}
