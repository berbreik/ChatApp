package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq" // Postgres driver
)

func main() {
	// Expect DATABASE_URL as first CLI argument
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run populatedb.go <DATABASE_URL>")
	}
	dsn := os.Args[1]

	// Connect to Postgres
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	// Load schema file
	schema, err := ioutil.ReadFile("internal/db/Schema.psql")
	if err != nil {
		log.Fatalf("failed to read schema: %v", err)
	}

	// Execute schema
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatalf("failed to apply schema: %v", err)
	}

	fmt.Println("✅ Database populated successfully")
}
