package entities

type AdminBase struct {
	Phone int64 `json:"phone"`
}

type AdminCreate struct {
	AdminBase
	Password string `json:"password"`
}

type AdminLogin struct {
	Phone    int64  `json:"phone"`
	Password string `json:"password"`
}

type Admin struct {
	AdminBase
	AdminID int `json:"id"`
}

type AdminChangePWD struct {
	AdminID int    `json:"id"`
	NewPWD  string `json:"newPWD"`
}
