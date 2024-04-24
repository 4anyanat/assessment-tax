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

	checknum := (reflect.TypeOf(taxin.TotalIncome).Kind() == reflect.String) || (reflect.TypeOf(taxin.Wht).Kind() == reflect.String) || (reflect.TypeOf(taxin.Allowances[0].Amount).Kind() == reflect.String)
	checkstr := (reflect.TypeOf(taxin.Allowances[0].AllowanceType).Kind() == reflect.String)

	if checknum && !checkstr {

		return c.JSON(http.StatusBadRequest, "InvalidInput: Invalid input types")
	}
	if taxin.TotalIncome < 0 || taxin.Wht < 0 || taxin.Allowances[0].Amount < 0 {

		return c.JSON(http.StatusBadRequest, "InvalidInput: Inputs are not positive")
	}

	taxCalced := taxCalc(taxin.TotalIncome)

	return c.JSON(http.StatusOK, map[string]float64{
		"tax": taxCalced,
	})
}

func taxCalc(totalIncome float64) float64 {

	taxrate := taxRate{
		rate35:      0.35,
		rate20:      0.20,
		rate15:      0.15,
		rate10:      0.10,
		rate0:       0.0,
		personalDec: 60000.0,
	}

	tax := 0.0
	stepVal := 0.0
	totalIncome -= float64(taxrate.personalDec)

	if totalIncome > 2000000 {
		stepVal = totalIncome - 2000000
		tax += stepVal * taxrate.rate35

		totalIncome -= stepVal
		stepVal = totalIncome - 1000000
		tax += stepVal * taxrate.rate20

		totalIncome -= stepVal
		stepVal = totalIncome - 500000
		tax += stepVal * taxrate.rate15

		totalIncome -= stepVal
		stepVal = totalIncome - 150000
		tax += stepVal * taxrate.rate10

	} else if totalIncome >= 1000001 && totalIncome <= 2000000 {
		stepVal = totalIncome - 1000000
		tax += stepVal * taxrate.rate20

		totalIncome -= stepVal
		stepVal = totalIncome - 500000
		tax += stepVal * taxrate.rate15

		totalIncome -= stepVal
		stepVal = totalIncome - 150000
		tax += stepVal * taxrate.rate10

	} else if totalIncome >= 500001 && totalIncome <= 1000000 {
		stepVal = totalIncome - 500000
		tax += stepVal * taxrate.rate15

		totalIncome -= stepVal
		stepVal = totalIncome - 150000
		tax += stepVal * taxrate.rate10

	} else if totalIncome >= 150001 && totalIncome <= 500000 {
		stepVal = totalIncome - 150000
		tax += stepVal * taxrate.rate10

	} else if totalIncome >= 0 && totalIncome <= 150000 {
		tax = taxrate.rate0
	}

	return tax
}
