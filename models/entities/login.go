package entities

import "gorm.io/gorm"

type SignUp struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
