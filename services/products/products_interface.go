package services

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/api"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
)

type Products interface {
	GetProducts(context.Context) ([]entities.Item, models.EcomError)
	GetUserDetails(context.Context) (entities.EcomUsers, models.EcomError)
	CardDetails(context.Context, models.CardDetails) (models.CardDetails, models.EcomError)
}

type products struct {
	APIProvider api.Service
	Store       repositories.RepositoryInterface
}

func NewService(apiProvider api.Service, store repositories.RepositoryInterface) Products {
	return &products{
		APIProvider: apiProvider,
		Store:       store,
	}
}
