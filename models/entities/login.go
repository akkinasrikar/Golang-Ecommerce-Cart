package entities

import "gorm.io/gorm"

type SignUp struct {
	gorm.Model
	UserId   int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:name;not null"`
	Email    string `gorm:"column:email;not null;unique"`
	Password string `gorm:"column:password;not null"`
}

type Login struct {
	Name     string `gorm:"column:name;not null"`
	Password string `gorm:"column:password;not null"`
}

func (SignUp) TableName() string {
	return "user_details"
}
