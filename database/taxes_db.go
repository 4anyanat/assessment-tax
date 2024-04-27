package database

import (
	"database/sql"
	"log"
	"os"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/labstack/echo/v4"
)

// Input and output of updated personal deduction amount
func TaxesUpdate(c echo.Context) error {
	url := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Database connection error", err)
	}
	defer db.Close()

	decinput := new(decInput)
	err = c.Bind(decinput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errMsg{Message: err.Error()})
	}

	var personalDec personalDec

	if decinput.Amount > 100000 || decinput.Amount < 10000 {
		return c.JSON(http.StatusBadRequest, "Personal deduction should be over 10,000 Bahts and not over than 100,000 Bahts")
	}else {

		stmt, err := db.Prepare("UPDATE taxes SET personalDeduction=$1;")
		
		if err != nil {
			log.Fatal("Prepare statement error:", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(decinput.Amount)
		if err != nil {
			log.Fatal("Update error", err)
		}

		personalDec.PersonalDeduction = decinput.Amount
	}
	
	return c.JSON(http.StatusOK, personalDec)
}

// Input and output of updated k-receipt allowance limit
func AllowanceUpdate(c echo.Context) error {
	url := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Database connection error", err)
	}
	defer db.Close()

	alwinput := new(decInput)
	err = c.Bind(alwinput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errMsg{Message: err.Error()})
	}

	var kReceipt kReceipt

	if alwinput.Amount > 100000 || alwinput.Amount < 0 {
		return c.JSON(http.StatusBadRequest, "k-receipt should be over 0 Bahts and not over than 100,000 Bahts")
	}else {

		stmt, err := db.Prepare("UPDATE taxes SET kReceipt=$1;")
		
		if err != nil {
			log.Fatal("Prepare statement error:", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(alwinput.Amount)
		if err != nil {
			log.Fatal("Update error", err)
		}

		kReceipt.KReceipt = alwinput.Amount
	}
	
	return c.JSON(http.StatusOK, kReceipt)
}