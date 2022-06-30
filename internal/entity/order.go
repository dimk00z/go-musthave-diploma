package entity

import "time"

const (
	OrderStatusNew        = "NEW"
	OrderStatusProcessing = "PROCESSING"
	OrderStatusRegistered = "REGISTERED"
	OrderStatusInvalid    = "INVALID"
)

type Order struct {
	OrderID     string    `json:"-"`
	OrderNumber string    `json:"number"`
	Status      string    `json:"status"`
	Accrual     float32   `json:"accrual"`
	ProcessedAt time.Time `json:"processed_at"`
	UserID      string    `json:"-"`
}
