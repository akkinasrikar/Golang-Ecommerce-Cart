package entities

type DeliveryAddress struct {
	AddressID string `gorm:"column:address_id;primaryKey"`
	HouseNo   string `gorm:"column:house_no;not null"`
	Street    string `gorm:"column:street;not null"`
	City      string `gorm:"column:city;not null"`
	State     string `gorm:"column:state;not null"`
	Pincode   string `gorm:"column:pincode;not null"`
	EcomID    string `gorm:"column:ecom_id;not null"`
}
