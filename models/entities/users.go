package entities

type EcomUsers struct {
	EcomID       string `gorm:"column:ecom_id;primaryKey"`
	EmailID      string `gorm:"column:email_id;not null"`
	AccountName  string `gorm:"column:account_name;not null"`
	WalletAmount int64  `gorm:"column:wallet_amount;not null;default:0"`
	UsersID      int64  `gorm:"column:users_id;not null"`
	CartItems    string `gorm:"column:cart_items"`
}
