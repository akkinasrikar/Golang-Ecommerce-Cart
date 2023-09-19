package models

type AddToCart struct {
	ProductId int    `json:"product_id" binding:"required"`
	Action    string `json:"action" binding:"required"`
}

type CartResponse struct {
	ProductID int    `json:"product_id"`
	Action    string `json:"action"`
	Message   string `json:"message"`
}
