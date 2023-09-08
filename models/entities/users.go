package entities

type EcomUsers struct {
	EcomID          string `gorm:"column:ecom_id;primaryKey"`
	AccountName     string `gorm:"column:account_name;not null"`
	WalletAmount    int64  `gorm:"column:wallet_amount;not null;default:0"`
	DeliveryAddress string `gorm:"column:delivery_address"`
	UsersID         int64  `gorm:"column:users_id;not null"`
}