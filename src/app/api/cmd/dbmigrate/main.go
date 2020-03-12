package main

import (
	"log"
	"os"

	"app/api/config"
	"app/api/pkg/logger"
)

func main() {
	// Create the logger.
	l := logger.New(log.New(os.Stderr, "", log.LstdFlags))

	// Migrate the database.
	config.Database(l)
	l.Printf("Database migration complete.")
}
