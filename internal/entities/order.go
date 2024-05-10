package entities

type OrderBase struct {
	UserID int   `json:"user_id"`
	Breads []int `json:"breads"`
}

type Order struct {
	OrderBase
	ID int `json:"id"`
}
