package main

import (
	"os"

	"github.com/nathanberry97/rss2go/src/routes"
	"github.com/nathanberry97/rss2go/src/services"
	"github.com/nathanberry97/rss2go/src/utils"
)

func main() {
	// Set environment variables
	utils.SetEnv(".env")
	port := os.Getenv("PORT")
	hostName := os.Getenv("HOST_NAME")

	// Set up the database
	db := services.DatabaseConnection()
	defer db.Close()

	err := services.InitializeDatabase(db)
	if err != nil {
		panic(err)
	}

	// Start the server
	router := routes.InitialiseRouter()
	router.Run(hostName + ":" + port)
}
