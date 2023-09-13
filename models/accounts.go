package models

type CardDetails struct {
	CardId     string `json:"card_id"`
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv"`
	Name       string `json:"name"`
	CardType   string `json:"card_type"`
}
