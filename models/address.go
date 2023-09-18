package models

type Address struct {
	HouseNo string `json:"house_no" binding:"required"`
	Street  string `json:"street" binding:"required"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	Pincode string `json:"pincode" binding:"required"`
}
