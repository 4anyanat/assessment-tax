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

    if err := os.MkdirAll("upload-csv", 0755); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create directory: "+err.Error())
	}

    file, err := c.FormFile("taxFile")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to retrieve tax file: "+err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open tax file: "+err.Error())
	}
	defer src.Close()

	// Create a .csv file
    dst, err := os.Create("upload-csv/" + file.Filename)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create file: "+err.Error())
    }
    defer dst.Close()

    if _, err = io.Copy(dst, src); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to copy file contents: "+err.Error())
    }

	csvFile, err := os.ReadFile("upload-csv/" + file.Filename)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read CSV file: "+err.Error())
    }

	taxContent := csv.NewReader(strings.NewReader(string(csvFile)))
	_, err = taxContent.Read()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to read CSV header: "+err.Error())
	}
	
    // Read .csv file
    for {
		row, err := taxContent.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to read CSV row: "+err.Error())
		}

		totalIncome, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid total income value: "+err.Error())
		}
		wht, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid WHT value: "+err.Error())
		}
		donation, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid donation value: "+err.Error())
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