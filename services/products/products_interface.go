package services

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/api"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
)

type Products interface {
	GetProducts(context.Context) (responses.ItemsResponse, models.EcomError)
}

type products struct {
	APIProvider api.Service
}

func NewService(apiProvider api.Service) Products {
	return &products{
		APIProvider: apiProvider,
	}
}
