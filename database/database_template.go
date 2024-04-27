package database

type personalDec struct {
	PersonalDeduction float64 `json:"personalDeduction"`
}

type kReceipt struct {
	KReceipt float64 `json:"kReceipt"`
}

type decInput struct {
	Amount float64 `json:"amount"`
}

type errMsg struct {
	Message string `json:"message"`
}