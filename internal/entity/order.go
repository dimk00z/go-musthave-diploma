package entity

type Order struct {
	OrderID     string  `json:"order"`
	Status      string  `json:"status"`
	Sum         float32 `json:"sum"`
	ProcessedAt string  `json:"processed_at"`
}
