package entity

type Withdrawal struct {
	OrderNumber string `json:"order"`
	Sum         string `json:"sum"`
	ProcessedAt string `json:"processed_at"`
}
