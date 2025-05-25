package database

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func DatabaseConnection() *sql.DB {
	conn, err := sql.Open("sqlite", "internal/database/rss.db")
	if err != nil {
		log.Fatalf("Unable to connect to SQLite database: %v\n", err)
	}

	conn.SetMaxOpenConns(1)
	conn.SetMaxIdleConns(1)

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping SQLite database: %v\n", err)
	}

	log.Println("Connected to SQLite database successfully.")
	return conn
}

func InitializeDatabase(db *sql.DB) error {
	_, err := db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return fmt.Errorf("error enabling foreign keys: %w", err)
	}

	file, err := os.Open("internal/database/init.sql")
	if err != nil {
		return fmt.Errorf("error reading init.sql: %w", err)
	}

	sqlBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading init.sql: %w", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("error executing init.sql: %w", err)
	}

	log.Println("Database initialized successfully.")
	return nil
}
