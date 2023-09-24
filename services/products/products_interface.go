package services

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/api"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
)

type Products interface {
	SeedData(context.Context) models.EcomError
	GetProducts(context.Context) ([]entities.Item, models.EcomError)
	GetProductById(context.Context, int) (string, models.EcomError)
	GetUserDetails(context.Context) (models.EcomUsers, models.EcomError)
	CardDetails(context.Context, models.CardDetails) (models.CardDetails, models.EcomError)
	GetCardDetails(context.Context) ([]models.CardDetails, models.EcomError)
	AddAddress(context.Context, models.Address) (entities.DeliveryAddress, models.EcomError)
	GetAddress(context.Context) ([]entities.DeliveryAddress, models.EcomError)
	AddOrDeleteToCart(context.Context, models.AddToCart) (models.CartResponse, models.EcomError)
	GetProductsFromCart(context.Context) ([]entities.Item, models.EcomError)
	OrderProducts(context.Context, models.PlaceOrder) (models.EcomOrderResponse, models.EcomError)
	GetOrdersByUserId(context.Context) ([]entities.Order, models.EcomError)
}

type products struct {
	APIProvider api.Service
	ProducAsynqService ProducAsynqService 
	Store       repositories.RepositoryInterface
}

func NewService(apiProvider api.Service, store repositories.RepositoryInterface, producAsynqService ProducAsynqService) Products {
	return &products{
		ProducAsynqService: producAsynqService,
		APIProvider: apiProvider,
		Store:       store,
	}
}
