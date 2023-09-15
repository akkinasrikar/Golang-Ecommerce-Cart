package models

type CardDetails struct {
	CardId     string `json:"card_id"`
	CardNumber int64  `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        int64  `json:"cvv"`
	Name       string `json:"name"`
	CardType   string `json:"card_type"`
}
