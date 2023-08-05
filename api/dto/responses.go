package dto

type Items struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

type ItemsResponse []Items
