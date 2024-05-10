package entities

type AdminBase struct {
	Phone int64 `json:"phone"`
}

type AdminCreate struct {
	AdminBase
	Password string `json:"password"`
}

type AdminLogin struct {
	Phone    int64  `json:"login"`
	Password string `json:"password"`
}

type Admin struct {
	AdminBase
	AdminID int `json:"id"`
}
