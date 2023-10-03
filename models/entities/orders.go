package entities

type Order struct {
	OrderID        string `gorm:"column:order_id;type:varchar(255);primary_key;not null"`
	ItemID         int    `gorm:"column:item_id;type:int4;not null"`
	OrderStatus    string `gorm:"column:order_status;type:varchar(255);not null"`
	OrderAmount    int64  `gorm:"column:order_amount;type:int8;not null"`
	OrderDate      string `gorm:"column:order_date;type:varchar(255);not null"`
	OrderName      string `gorm:"column:order_name;type:varchar(5000);not null"`
	PaymentMode    string `gorm:"column:payment_mode;type:varchar(255);not null"`
	DeliveryStatus string `gorm:"column:delivery_status;type:varchar(255);not null"`
	DeliveryDate   string `gorm:"column:delivery_date;type:varchar(255);not null"`
	AddressID      string `gorm:"column:address_id;type:varchar(500);not null"`
	CardID         string `gorm:"column:card_id;type:varchar(255);not null"`
	EcomID         string `gorm:"column:ecom_id;type:varchar(255);not null"`
	UsersID        int64  `gorm:"column:users_id;type:int8;not null"`
}
