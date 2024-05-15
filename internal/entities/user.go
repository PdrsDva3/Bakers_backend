package entities

type UserBase struct {
	Phone int64  `json:"phone"`
	Name  string `json:"name"`
}

type UserCreate struct {
	UserBase
	Password string `json:"password"`
}

type UserLogin struct {
	Phone    int64  `json:"phone"`
	Password string `json:"password"`
}

type UserChangePassword struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type UserChangeName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	UserBase
	ID int `json:"id"`
}
