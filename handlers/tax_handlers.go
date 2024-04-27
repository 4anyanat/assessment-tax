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

type taxRate struct {
	rate35      float64
	rate20      float64
	rate15      float64
	rate10      float64
	rate0       float64
	personalDec float64
}

func Tax_Cal_Handler(c echo.Context) error {
	taxin := new(taxInput)
	err := c.Bind(taxin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errMsg{Message: err.Error()})
	}

	if (taxin.TotalIncome < 0) || (taxin.Wht < 0) || (taxin.Wht > taxin.TotalIncome){

		return c.JSON(http.StatusBadRequest, "InvalidInput: Inputs are incorrect")
	}

	// type TaxCalculator func(totalIncome float64, wht float64, allowances float64) taxCalcInfo

	// var calculateTax TaxCalculator = taxCalc

	// totalAllowance := checkAllowance(taxin.Allowances)

	// // Execute tax calculation
	// taxCalced := calculateTax(taxin.TotalIncome, taxin.Wht, totalAllowance)

	taxCalced := CalculateTax(*taxin)

	return c.JSON(http.StatusOK, taxCalced)
}

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

func taxCalc(totalIncome float64, wht float64, allowances float64) taxCalcInfo {
	// Declare tax level rates
	taxrate := taxRate{
		rate35:      0.35,
		rate20:      0.20,
		rate15:      0.15,
		rate10:      0.10,
		rate0:       0.0,
		personalDec: 60000.0,
	}

	// Declare taxes for each level
	var (
		taxLvl35 float64
		taxLvl20 float64
		taxLvl15 float64
		taxLvl10 float64
		taxLvl0  float64
	)

	var taxOutput taxCalcInfo
	var taxType   string
	var refund	  float64

	// Initialization
	totalTax := 0.0
	totalIncome -= float64(taxrate.personalDec)
	totalIncome -= allowances

	// Calculate tax for each level
	if totalIncome > 2000000 {
		stepVal := totalIncome - 2000000
		taxLvl35 = stepVal * taxrate.rate35
		totalIncome -= stepVal
	}
	
	if totalIncome > 1000000 {
		stepVal := totalIncome - 1000000
		taxLvl20 = stepVal * taxrate.rate20
		totalIncome -= stepVal
	}

	if totalIncome > 500000 {
		stepVal := totalIncome - 500000
		taxLvl15 = stepVal * taxrate.rate15
		totalIncome -= stepVal
	}
	
	if totalIncome > 150000 {
		stepVal := totalIncome - 150000
		taxLvl10 = stepVal * taxrate.rate10
		totalIncome -= stepVal
	}
	
	if totalIncome > 0 {
		taxLvl0 = taxrate.rate0
	}

	totalTax = taxLvl35 + taxLvl20 + taxLvl15 + taxLvl10 + taxLvl0

	// Calcaulate the total tax
	if totalTax >= wht {
		totalTax -= wht
		refund = 0
		taxType = "tax"
	}else {
		refund = wht - totalTax
		taxType = "refund"
	}

	// Output tax for all levels
	taxOutput.Tax = totalTax
	taxOutput.TaxLvl35 = taxLvl35
	taxOutput.TaxLvl20 = taxLvl20
	taxOutput.TaxLvl15 = taxLvl15
	taxOutput.TaxLvl10 = taxLvl10
	taxOutput.TaxLvl0 = taxLvl0
	taxOutput.Refund = refund
	taxOutput.TaxType = taxType

	return taxOutput
}

func CalculateTax(taxin taxInput) TaxOutput {
	var taxOutput taxOutputInfo
	var taxRefundInfo taxRefundInfo

	totalAllowance := checkAllowance(taxin.Allowances)
	
	// Execute tax calculation
	taxCalced := taxCalc(taxin.TotalIncome, taxin.Wht, totalAllowance)
	
	if taxCalced.TaxType == "tax" {
		taxOutput = taxLvlOutput(taxCalced)
		return taxOutput
	} else if taxCalced.TaxType == "refund" {
		taxRefundInfo = taxRefundOutput(taxCalced)
		return taxRefundInfo
	}

	return taxOutput
}

func (t taxOutputInfo) GetOutput() interface{} {
	return t
}

func (t taxRefundInfo) GetOutput() interface{} {
	return t
}

func taxLvlOutput(taxInfo taxCalcInfo) taxOutputInfo {
	var taxOutput taxOutputInfo

	taxOutput.Tax = taxInfo.Tax

	taxOutput.TaxLevel = make([]taxLevel, 5)
	taxOutput.TaxLevel[0].Tax = taxInfo.TaxLvl0
	taxOutput.TaxLevel[1].Tax = taxInfo.TaxLvl10
	taxOutput.TaxLevel[2].Tax = taxInfo.TaxLvl15
	taxOutput.TaxLevel[3].Tax = taxInfo.TaxLvl20
	taxOutput.TaxLevel[4].Tax = taxInfo.TaxLvl35

	taxOutput.TaxLevel[0].Level = "0-150,000"
	taxOutput.TaxLevel[1].Level = "150,001-500,000"
	taxOutput.TaxLevel[2].Level = "500,001-1,000,000"
	taxOutput.TaxLevel[3].Level = "1,000,001-2,000,000"
	taxOutput.TaxLevel[4].Level = "2,000,001 ขึ้นไป"
	
	return taxOutput
}

func taxRefundOutput(taxInfo taxCalcInfo) taxRefundInfo {
	var taxOutput taxRefundInfo

	taxOutput.TaxRefund = taxInfo.Refund

	taxOutput.TaxLevel = make([]taxLevel, 5)
	taxOutput.TaxLevel[0].Tax = taxInfo.TaxLvl0
	taxOutput.TaxLevel[1].Tax = taxInfo.TaxLvl10
	taxOutput.TaxLevel[2].Tax = taxInfo.TaxLvl15
	taxOutput.TaxLevel[3].Tax = taxInfo.TaxLvl20
	taxOutput.TaxLevel[4].Tax = taxInfo.TaxLvl35

	taxOutput.TaxLevel[0].Level = "0-150,000"
	taxOutput.TaxLevel[1].Level = "150,001-500,000"
	taxOutput.TaxLevel[2].Level = "500,001-1,000,000"
	taxOutput.TaxLevel[3].Level = "1,000,001-2,000,000"
	taxOutput.TaxLevel[4].Level = "2,000,001 ขึ้นไป"

	return taxOutput
}

func checkAllowance(allowances []allowance) float64 {
	var donationAll float64
	
	for _, allowance := range allowances {
		// Donation limit setting >> Max. 10,000 Bahts
		if allowance.AllowanceType == "donation"{
			if allowance.Amount >= 100000{
				donationAll = 100000.0
				return donationAll
			}else if allowance.Amount < 100000 && allowance.Amount >= 0{
				donationAll = allowance.Amount
				return donationAll
			}
		}
	}
	return 0
}