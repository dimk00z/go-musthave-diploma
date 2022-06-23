package entity

type Balance struct {
	Current float32 `json:"current"`
	Spend   float32 `json:"withdrawn"`
}
