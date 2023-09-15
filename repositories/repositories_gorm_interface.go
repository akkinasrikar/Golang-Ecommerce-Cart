package repositories

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/database"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
)

type Repository struct {
	dbStore database.DB
}

func NewRepository(dbStore database.DB) RepositoryInterface {
	return &Repository{
		dbStore: dbStore,
	}
}

//go:generate mockgen -package mocks -source=repositories_gorm_interface.go -destination=mocks/repositories_gorm_interface_mocks.go
type RepositoryInterface interface {
	SignUp(userDetails entities.SignUp) (entities.SignUp, models.EcomError)
	Login(userDetails entities.Login) (entities.SignUp, models.EcomError)
	GetAllProducts() ([]entities.Item, models.EcomError)
	CreateEcomAccount(ecomAccountDetails entities.EcomUsers) (entities.EcomUsers, models.EcomError)
	GetUserDetails(ctx context.Context) (entities.EcomUsers, models.EcomError)
	CreateCardDetails(cardDetails entities.CardDetails) (entities.CardDetails, models.EcomError)
	GetCardDetails(userDetails entities.EcomUsers) ([]entities.CardDetails, models.EcomError)
}
