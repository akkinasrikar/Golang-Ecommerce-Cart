package models

type EcomUsers struct {
	EcomID          string `json:"ecom_id"`
	AccountName     string `json:"account_name"`
	WalletAmount    int64  `json:"wallet_amount"`
	DeliveryAddress string `json:"delivery_address"`
	UsersID         int64  `json:"users_id"`
	CartItems       []int  `json:"cart_items"`
}
