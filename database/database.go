package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func init(){
	fmt.Println("Initializing")
}

func DatabaseInit() {
	// Get database url from environment variable DATABASE_URL
	url := os.Getenv("DATABASE_URL")
	if url != "" {
		fmt.Println("DATABASE_URL is set to:", url)
	} else {
		fmt.Println("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", url)
	if err != nil{
		log.Fatal("Database connection error", err)
	}
	defer db.Close()

	createTableIfNotExists(db)

	if !hasRows(db, "taxes") {
		// Insert default values if no rows are present
		insertDefaults(db)
	}
	fmt.Println("Database initialization successful")
}

func createTableIfNotExists(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS taxes (
		personalDeduction FLOAT, 
		kReceipt FLOAT
	)`
	if _, err := db.Exec(query); err != nil {
		log.Fatal("Table creation error", err)
	}
}

func hasRows(db *sql.DB, tableName string) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s LIMIT 1)", tableName)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		log.Fatal("Error checking for rows", err)
	}
	return exists
}

func insertDefaults(db *sql.DB) {
	query := "INSERT INTO taxes (personalDeduction, kReceipt) VALUES ($1, $2)"
	_, err := db.Exec(query, 60000.0, 50000.0)
	if err != nil {
		log.Fatal("Error inserting default values", err)
	}
	fmt.Println("Inserted default values")
}