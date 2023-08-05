package repositories

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) SignUp(userDetails entities.SignUp) (entities.SignUp, models.EcomError) {
	if err := r.dbStore.Create(&userDetails).Error; err != nil {
		return entities.SignUp{}, *helper.ErrorInternalSystemError("Error while signing up : " + err())
	}
	return userDetails, models.EcomError{}
}

func (r *Repository) Login(userDetails entities.Login) (entities.Login, models.EcomError) {
	var user entities.Login
	if err := r.dbStore.Where("name = ? OR email = ?", userDetails.Name, userDetails.Name).First(&user).Error; err != nil {
		return entities.Login{}, *helper.ErrorInternalSystemError("Error while logging in : " + err())
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))
	if err != nil {
		return entities.Login{}, *helper.ErrorInternalSystemError("Error while comparing password : " + err.Error())
	}
	return user, models.EcomError{}
}
