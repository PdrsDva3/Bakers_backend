package entities

type BreadBase struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Count       int64   `json:"count"`
	Photo       string  `json:"photo"`
}

type Bread struct {
	BreadBase
	ID int `json:"id"`
}
