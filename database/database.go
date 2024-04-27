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

	createDB := `
	CREATE TABLE IF NOT EXISTS taxes ( personalDeduction FLOAT )
	`

	_, err = db.Exec(createDB)
	if err != nil{
		log.Fatal("Table creation error", err)
	}

	fmt.Println("Successful")

	// var personalDeduction = 60000.0

	// row := db.QueryRow("INSERT INTO taxes (personalDeduction) values ($1) RETURNING personalDeduction", 60000.0)
	// // db.Exec("DELETE FROM taxes")
	// err = row.Scan(&personalDeduction)
	// if err != nil {
	// 	log.Fatal("Error", err)
	// }
}