package models

type OrderDetailsEmail struct {
	OrderID string `json:"order_id"`
}

type SendEmailRequest struct {
	Email       string `json:"email"`
	Message     string `json:"message"`
	Subject     string `json:"subject"`
	ImageBase64 string `json:"image_base64"`
}
