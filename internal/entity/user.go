package entity

type User struct {
	UserId  string `json:"user_id"`
	Login   string `json:"user"`
	Balance Balance
	Orders  []Order
}
