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

type errMsg struct {
	Message string `json:"message"`
}
