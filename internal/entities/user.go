package entities

type UserBase struct {
	Phone int64  `json:"phone"`
	Name  string `json:"name"`
}

type UserCreate struct {
	UserBase
	Password string `json:"password"`
}

type User struct {
	UserBase
	ID int `json:"id"`
}
