package models

type EcomOrders struct {
	IsFromWallet bool   `json:"is_from_wallet"`
	AddressID    string `json:"address_id" binding:"required"`
	CardId       string `json:"card_id" binding:"required"`
}

type PlaceOrder struct {
	EcomId     string `json:"ecom_id"`
	UsersID    int64  `json:"users_id"`
	AddressID  string `json:"address_id"`
	Address    string `json:"address"`
	CardNumber int64 `json:"card_number"`
	CardId     string `json:"card_id"`
}

type EcomOrderResponse struct {
	Orders  []OrderDetails `json:"orders"`
	Message string         `json:"message"`
}

type OrderDetails struct {
	OrderID      string `json:"order_id"`
	Amount       int64  `json:"amount"`
	ProductName  string `json:"product_name"`
	OrderedDate  string `json:"ordered_date"`
	Address      string `json:"address"`
	DeliveryDate string `json:"delivery_date"`
	CardNUmber  string  `json:"card_number"`
}
