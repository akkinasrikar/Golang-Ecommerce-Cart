package entities

type CardDetails struct {
	EncryptedData string `gorm:"column:encrypted_data;not null"`
	EcomId        string `gorm:"column:ecom_id;not null"`
}
