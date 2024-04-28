package handlers

import (
	"testing"
    _ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCheckAllowance(t *testing.T) {

	allowances := []allowance{
		{AllowanceType: "k-receipt", Amount: 200000.0},
		{AllowanceType: "donation", Amount: 100000.0},
	}

	result := checkAllowance(allowances)

	expectedOutput := 150000.0

	assert.Equal(t, expectedOutput, result)
}

func TestTaxCalc(t *testing.T) {

	totalIncome := 500000.0
	wht := 0.0
	allowances := 150000.0

	result := taxCalc(totalIncome, wht, allowances)

	expectedOutput := taxCalcInfo{
		Tax: 14000.0,
		TaxLvl35: 0.0,
		TaxLvl20: 0.0,
		TaxLvl15: 0.0,
		TaxLvl10: 14000.0,
		TaxLvl0:  0.0,
		Refund:   0.0,
		TaxType:   "tax",
	}

	assert.Equal(t, expectedOutput, result)
}
