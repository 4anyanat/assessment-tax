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

	// Create table taxes if not already exists
	createDB := `
	CREATE TABLE IF NOT EXISTS taxes ( personalDeduction FLOAT, kReceipt FLOAT )
	`

	_, err = db.Exec(createDB)
	if err != nil{
		log.Fatal("Table creation error", err)
	}

	// // Insert new row of table (personalDeduction, kReceipt)
	// stmt, err := db.Prepare("INSERT INTO taxes (personalDeduction, kReceipt) VALUES ($1, $2)")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// defer stmt.Close()

	// row, _ := stmt.Exec(60000.0, 0.0)
	// fmt.Println(row.RowsAffected())

	fmt.Println("Successful")
}