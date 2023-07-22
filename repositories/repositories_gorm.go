package repositories

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) SignUp(userDetails entities.SignUp) (entities.SignUp, models.EcomError) {
	err := r.Db.Create(&userDetails).Error
	if err != nil {
		return entities.SignUp{}, *helper.ErrorInternalSystemError("Error while signing up : " + err.Error())
	}
	return userDetails, models.EcomError{}
}

func (r *Repository) Login(userDetails entities.Login) (entities.Login, models.EcomError) {
	var user entities.Login
	err := r.Db.Table("user_details").Where("name = ? OR email = ?", userDetails.Name, userDetails.Name).First(&user).Error
	if err != nil {
		return entities.Login{}, *helper.ErrorInternalSystemError("Error while logging in : " + err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))
	if err != nil {
		return entities.Login{}, *helper.ErrorInternalSystemError("Error while comparing password : " + err.Error())
	}
	return user, models.EcomError{}
}
