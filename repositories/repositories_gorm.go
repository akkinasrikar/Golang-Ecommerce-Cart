package repositories

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) SignUp(userDetails entities.SignUp) (entities.SignUp, error) {
	err := r.Db.Create(&userDetails).Error
	if err != nil {
		return entities.SignUp{}, err
	}
	return userDetails, nil
}

func (r *Repository) Login(userDetails entities.Login) (entities.Login, error) {
	var user entities.Login
	err := r.Db.Table("user_details").Where("user_name = ? OR user_email = ?", userDetails.Name, userDetails.Name).First(&user).Error
	if err != nil {
		return entities.Login{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))
	if err != nil {
		return entities.Login{}, err
	}
	return user, nil
}
