package handlers

import (
	"net/http"
	"reflect"

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

	checknum := (reflect.TypeOf(taxin.TotalIncome).Kind() == reflect.String) || (reflect.TypeOf(taxin.Wht).Kind() == reflect.String)

	var invalidAllowance bool
	var invalidAllowanceTypeof bool

	for _, allowance := range taxin.Allowances{
		if allowance.Amount < 0{
			invalidAllowance = true
		}
		if reflect.TypeOf(allowance.Amount).Kind() == reflect.String{
			invalidAllowanceTypeof = true
		}else if reflect.TypeOf(allowance.AllowanceType).Kind() != reflect.String{
			invalidAllowanceTypeof = true
		}
	}

	if checknum || invalidAllowanceTypeof{

		return c.JSON(http.StatusBadRequest, "InvalidInput: Invalid input types")
	}
	if (taxin.TotalIncome < 0) || (taxin.Wht < 0) || (taxin.Wht > taxin.TotalIncome) || invalidAllowance {

		return c.JSON(http.StatusBadRequest, "InvalidInput: Inputs are incorrect")
	}

	taxCalced := taxCalc(taxin.TotalIncome, taxin.Wht, taxin.Allowances)

	return c.JSON(http.StatusOK, taxCalced)
}

func taxCalc(totalIncome float64, wht float64, allowances []allowance) taxOutput {

	taxrate := taxRate{
		rate35:      0.35,
		rate20:      0.20,
		rate15:      0.15,
		rate10:      0.10,
		rate0:       0.0,
		personalDec: 60000.0,
	}

	var (
		taxLvl35 float64
		taxLvl20 float64
		taxLvl15 float64
		taxLvl10 float64
	)

	var taxOutput taxOutput

	tax := 0.0
	totalTax := 0.0
	totalIncome -= float64(taxrate.personalDec)

	for _, allowance := range allowances {
		if allowance.AllowanceType == "donation"{
			if allowance.Amount > 100000{
				totalIncome -= 100000
			}else{
				totalIncome -= allowance.Amount
			}
		}
	}

	if totalIncome > 2000000 {
		stepVal := totalIncome - 2000000
		taxLvl35 = stepVal * taxrate.rate35
		tax += taxLvl35
		totalIncome -= stepVal
		totalTax += tax
	}
	
	if totalIncome > 1000000 {
		stepVal := totalIncome - 1000000
		taxLvl20 = stepVal * taxrate.rate20
		tax += taxLvl20
		totalIncome -= stepVal
		totalTax += tax
	}

	if totalIncome > 500000 {
		stepVal := totalIncome - 500000
		taxLvl15 = stepVal * taxrate.rate15
		tax += taxLvl15
		totalIncome -= stepVal
		totalTax += tax
	}
	
	if totalIncome > 150000 {
		stepVal := totalIncome - 150000
		taxLvl10 = stepVal * taxrate.rate10
		tax += taxLvl10
		totalIncome -= stepVal
		totalTax += tax
	}
	if totalIncome > 0 {
		tax = taxrate.rate0
		totalTax += tax
	}

	totalTax -= wht

	taxOutput.Tax = totalTax

	taxOutput.TaxLevel = make([]taxLevel, 5)
	taxOutput.TaxLevel[0].Tax = taxrate.rate0
	taxOutput.TaxLevel[1].Tax = taxLvl10
	taxOutput.TaxLevel[2].Tax = taxLvl15
	taxOutput.TaxLevel[3].Tax = taxLvl20
	taxOutput.TaxLevel[4].Tax = taxLvl35

	taxOutput.TaxLevel[0].Level = "0-150,000"
	taxOutput.TaxLevel[1].Level = "150,001-500,000"
	taxOutput.TaxLevel[2].Level = "500,001-1,000,000"
	taxOutput.TaxLevel[3].Level = "1,000,001-2,000,000"
	taxOutput.TaxLevel[4].Level = "2,000,001 ขึ้นไป"

	return taxOutput
}

