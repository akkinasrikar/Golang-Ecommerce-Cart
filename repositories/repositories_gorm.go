package repositories

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) SignUp(userDetails entities.SignUp) (entities.SignUp, models.EcomError) {
	if err := r.dbStore.Create(&userDetails); err != nil {
		return entities.SignUp{}, *helper.ErrorInternalSystemError("Error while signing up : " + err.Error())
	}
	return userDetails, models.EcomError{}
}

func (r *Repository) Login(userDetails entities.Login) (entities.Login, models.EcomError) {
	var user entities.SignUp
	_, err := r.dbStore.Where("user_name = ? OR user_email = ?", userDetails.Name, userDetails.Name).Find(&user)
	if err != nil {
		return entities.Login{}, *helper.ErrorInternalSystemError(err.Error())
	}
	if user.Name == "" {
		return entities.Login{}, *helper.ErrorInternalSystemError("User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))
	if err != nil {
		return entities.Login{}, *helper.ErrorInternalSystemError("Password is incorrect")
	}
	return userDetails, models.EcomError{}
}
