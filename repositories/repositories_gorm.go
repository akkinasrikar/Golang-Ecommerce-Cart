package repositories

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
)

func (r *Repository) SignUp(userDetails entities.SignUp) (entities.SignUp, error) {
	err := r.Db.Create(&userDetails).Error
	if err != nil {
		return entities.SignUp{}, err
	}
	return userDetails, nil
}

// Login
func (r *Repository) Login(userDetails entities.Login) (entities.Login, error) {
	var user entities.Login
	err := r.Db.Table("user_details").Where("user_name = ? OR user_email = ? and user_password = ?", userDetails.Name, userDetails.Name, userDetails.Password).First(&user).Error
	if err != nil {
		return entities.Login{}, err
	}
	return user, nil
}
