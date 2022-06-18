package entity

type User struct {
	Login    string `json:"user"`
	Password string `json:"password"`
	Balance  Balance
	Orders   []Order
}
