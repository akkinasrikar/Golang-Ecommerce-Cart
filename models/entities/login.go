package entities

import "gorm.io/gorm"

type SignUp struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:user_name;not null"`
	Email    string `gorm:"column:user_email;not null;unique"`
	Password string `gorm:"column:user_password;not null"`
}

type Login struct {
	Name     string `gorm:"column:user_name;not null"`
	Password string `gorm:"column:user_password;not null"`
}

func (SignUp) TableName() string {
	return "user_details"
}
