package entities

type CardDetails struct {
	CardID        string `gorm:"column:card_id;primaryKey"`
	EncryptedData string `gorm:"column:encrypted_data;not null"`
	EcomId        string `gorm:"column:ecom_id;not null"`
}
