package models

type PayslipItem struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type PayslipResponse struct {
	Period          string        `json:"period"`
	Earnings        []PayslipItem `json:"earnings"`
	Deductions      []PayslipItem `json:"deductions"`
	TotalEarnings   float64       `json:"total_earnings"`
	TotalDeductions float64       `json:"total_deductions"`
	NetSalary       float64       `json:"net_salary"`
}
