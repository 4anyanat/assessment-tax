package database

type personalDec struct {
	PersonalDeduction float64 `json:"personalDeduction"`
}

type decInput struct {
	Amount float64 `json:"amount"`
}

type errMsg struct {
	Message string `json:"message"`
}