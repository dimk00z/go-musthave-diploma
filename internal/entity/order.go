package entity

type Order struct {
	OrderID     string  `json:"order_id"`
	OrderNumber string  `json:"number"`
	Status      string  `json:"status"`
	Accrual     float32 `json:"accrual"`
	ProcessedAt int64   `json:"processed_at"`
}
