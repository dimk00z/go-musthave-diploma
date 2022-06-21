package entity

type User struct {
	UserID   string `json:"user_id"`
	Login    string `json:"user"`
	Password string `json:"-"`
	Balance  Balance
	Orders   []Order
}
