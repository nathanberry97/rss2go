package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func DatabaseConnection() *pgx.Conn {
	connURL := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), connURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!")

	return conn
}
