package entity

type Order struct {
	Order_id     string  `json:"order"`
	Status       string  `json:"status"`
	Sum          float32 `json:"sum"`
	Processed_at string  `json:"processed_at"`
}
