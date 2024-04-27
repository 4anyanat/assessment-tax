package handlers

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func Tax_Csv_Handler(c echo.Context) error {

	var taxValues []taxCsvInput
	var taxInfoSlice []taxesInfo

    file, err := c.FormFile("csvFile")
    if err != nil {
        return err
    }

    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

	// Create a .csv file
    dst, err := os.Create("upload-csv/" + file.Filename)
    if err != nil {
        return err
    }
    defer dst.Close()

    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

	csvFile, err := os.ReadFile("upload-csv/" + file.Filename)
    if err != nil {
        return err
    }

	taxContent := csv.NewReader(strings.NewReader(string(csvFile)))
	_, err = taxContent.Read()
	if err != nil {
		return err
	}
	
    // Read .csv file
    for {
        row, err := taxContent.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }

        totalIncome, err := strconv.ParseFloat(row[0], 64)
        if err != nil {
            return err
        }
        wht, err := strconv.ParseFloat(row[1], 64)
        if err != nil {
            return err
        }
        donation, err := strconv.ParseFloat(row[2], 64)
        if err != nil {
            return err
        }

        taxValues = append(taxValues, taxCsvInput{
            TotalIncome: totalIncome,
            Wht:         wht,
            Donation:    donation,
        })
    }

	
	for _, taxValue := range taxValues {
		taxCalcued := taxCalc(taxValue.TotalIncome, taxValue.Wht, taxValue.Donation)

		taxInfo := taxesInfo{
			TotalIncome: taxValue.TotalIncome,
			Tax:         taxCalcued.Tax,
		}

		taxInfoSlice = append(taxInfoSlice, taxInfo)
	}
	
	taxData := taxes{
    	Taxes: taxInfoSlice,
	}
    return c.JSON(http.StatusOK, taxData)
}