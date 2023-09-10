package repositories

import (
	"context"

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

func (r *Repository) Login(userDetails entities.Login) (entities.SignUp, models.EcomError) {
	var user entities.SignUp
	_, err := r.dbStore.Where("user_name = ? OR user_email = ?", userDetails.Name, userDetails.Name).Find(&user)
	if err != nil {
		return entities.SignUp{}, *helper.ErrorInternalSystemError(err.Error())
	}
	if user.Name == "" {
		return entities.SignUp{}, *helper.ErrorInternalSystemError("User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))
	if err != nil {
		return entities.SignUp{}, *helper.ErrorInternalSystemError("Password is incorrect")
	}
	return user, models.EcomError{}
}

func (r *Repository) GetAllProducts() ([]entities.Item, models.EcomError) {
	var items []entities.Item
	_, err := r.dbStore.Find(&items)
	if err != nil {
		return []entities.Item{}, *helper.ErrorInternalSystemError(err.Error())
	}
	return items, models.EcomError{}
}

func (r *Repository) CreateEcomAccount(ecomAccountDetails entities.EcomUsers) (entities.EcomUsers, models.EcomError) {
	if err := r.dbStore.Create(&ecomAccountDetails); err != nil {
		return entities.EcomUsers{}, *helper.ErrorInternalSystemError("Error while creating ecom account : " + err.Error())
	}
	return ecomAccountDetails, models.EcomError{}
}

func (r *Repository) GetUserDetails(ecomCtx context.Context) (entities.EcomUsers, models.EcomError) {
	var user entities.EcomUsers
	authData := ecomCtx.Value(models.EcomctxKey("AuthData")).(models.AuthData)
	_, err := r.dbStore.Where("users_id = ?", authData.UsersId).Find(&user)
	if err != nil {
		return entities.EcomUsers{}, *helper.ErrorInternalSystemError(err.Error())
	}
	return user, models.EcomError{}
}
