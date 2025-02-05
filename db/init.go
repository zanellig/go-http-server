package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:redacted@(localhost:3306)/"

func Init() {
	file, err := os.Open("db/init.sql")
	if err != nil {
		log.Fatalf("Failed to open SQL file: %v", err)
	}
	defer file.Close()

	script, err := io.ReadAll(file) 
	if err != nil {
		log.Fatalf("Error reading init.sql: %v", err)
	}

	queries := strings.Split(string(script), ";")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error executing query: %s\n%v", query, err)
		}
	}

	fmt.Println("Database initialized")

	
}