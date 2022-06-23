package entity

type User struct {
	UserID   string  `json:"user_id"`
	Login    string  `json:"user"`
	Password string  `json:"-"`
	Orders   []Order `json:"orders"`
	Balance  Balance `json:"balance"`
}
