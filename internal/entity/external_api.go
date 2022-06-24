package entity

type AccrualSystemRequest struct {
	OrderNumber string `json:"number"`
}

type AccrualSystemResponse struct {
	OrderNumber string  `json:"number"`
	OrderStatus string  `json:"status"`
	Accrual     float32 `json:"accrual"`
}
