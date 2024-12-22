package main

import (
	"github.com/nathanberry97/rss2go/src/routes"
	"github.com/nathanberry97/rss2go/src/utils"
	"os"
)

func main() {
	// Set environment variables
	utils.SetEnv(".env")
	port := os.Getenv("PORT")
	hostName := os.Getenv("HOST_NAME")

	// Start the server
	router := routes.InitialiseRouter()
	router.Run(hostName + ":" + port)
}
