package main

import (
	"os"

	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/utils"
	"github.com/nathanberry97/rss2go/pkg/routes"
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
		panic(err)
	}

	// Start the server
	router := routes.InitialiseRouter()
	router.Run(hostName + ":" + port)
}
