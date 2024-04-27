package handlers

type taxInput struct {
	TotalIncome float64     `json:"totalIncome"`
	Wht         float64     `json:"wht"`
	Allowances  []allowance `json:"allowances"`
}

type allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type taxOutputInfo struct {
	Tax 		float64 `json:"tax"`
	TaxLevel []taxLevel `json:"taxLevel"`
}

type taxRefundInfo struct {
	TaxRefund float64 	`json:"taxRefund"`
	TaxLevel []taxLevel `json:"taxLevel"`
}

type taxLevel struct {
	Level 	string 	`json:"level"`
	Tax		float64 `json:"tax"`
}

type taxCsvInput struct {
	TotalIncome float64 `json:"totalIncome"`
	Wht         float64 `json:"wht"`
	Donation  	float64 `json:"donation"`
}
type taxes struct {
	Taxes []taxesInfo `json:"taxes"`
}

type taxCalcInfo struct {
	Tax 	 float64 
	TaxLvl35 float64
	TaxLvl20 float64
	TaxLvl15 float64
	TaxLvl10 float64
	TaxLvl0  float64
	Refund   float64
	TaxType	 string
}

type taxesInfo struct {
	TotalIncome float64 `json:"totalIncome"`
	Tax 		float64 `json:"tax"`
}

type TaxOutput interface {
	GetOutput() interface{}
}

type errMsg struct {
	Message string `json:"message"`
}
