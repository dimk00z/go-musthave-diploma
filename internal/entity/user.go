package entity

type User struct {
	UserID   string  `json:"user_id"`
	Login    string  `json:"user"`
	Password string  `json:"-"`
	Balance  float32 `json:"balance"`
	Spend    float32 `json:"spend"`
	Orders   []Order `json:"orders"`
}
