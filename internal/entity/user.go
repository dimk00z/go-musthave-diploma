package entity

type User struct {
	UserId   string `json:"user_id"`
	Login    string `json:"user"`
	Password string `json:"-"`
	Balance  Balance
	Orders   []Order
}
