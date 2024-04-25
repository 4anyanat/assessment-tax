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

type taxOutput struct {
	Tax 		float64 `json:"tax"`
	TaxLevel []taxLevel `json:"taxLevel"`
}

type taxLevel struct {
	Level 	string 	`json:"level"`
	Tax		float64 `json:"tax"`
}

type errMsg struct {
	Message string `json:"message"`
}
