package repositories

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{
		Db: db,
	}
}

type RepositoryInterface interface {
	SignUp(userDetails entities.SignUp) (entities.SignUp, models.EcomError)
	Login(userDetails entities.Login) (entities.Login, models.EcomError)
}
