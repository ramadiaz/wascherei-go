package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type MonthlyIncome struct {
	Year    int `json:"year"`
	Month   int `json:"month"`
	Total   int `json:"total"`
	Average int `json:"average"`
}

type IncomeResult struct {
	TotalIncome     int64
	TransactionCount int64
}