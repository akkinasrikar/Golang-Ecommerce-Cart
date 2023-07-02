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
