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
	GetProductById(id int) (entities.Item, models.EcomError)
	CreateProduct(item entities.Item) (entities.Item, models.EcomError)
	UpdateProductByID(id int, item entities.Item) (entities.Item, models.EcomError)
	CreateEcomAccount(ecomAccountDetails entities.EcomUsers) (entities.EcomUsers, models.EcomError)
	UpdateEcomAccount(ecomAccountDetails entities.EcomUsers, ecomId string) (entities.EcomUsers, models.EcomError)
	GetUserDetails(ctx context.Context) (entities.EcomUsers, models.EcomError)
	GetUserDetailsById(usersId int64) (entities.EcomUsers, models.EcomError)
	CreateCardDetails(cardDetails entities.CardDetails) (entities.CardDetails, models.EcomError)
	GetCardDetails(userDetails entities.EcomUsers) ([]entities.CardDetails, models.EcomError)
	CreateAddress(addressDetails entities.DeliveryAddress) (entities.DeliveryAddress, models.EcomError)
	GetAddress(userDetails entities.EcomUsers) ([]entities.DeliveryAddress, models.EcomError)
	GetAddressById(addressId string) (entities.DeliveryAddress, models.EcomError)
	AddToCart(userDetails entities.EcomUsers, Id int) (entities.Item, models.EcomError)
	GetProductFromCart(itemId int) (entities.Item, models.EcomError)
	GetCardDetailsById(cardId string) (entities.CardDetails, models.EcomError)
	CreateOrder(orderDetails entities.Order) (entities.Order, models.EcomError)
	UpdateOrderByID(orderId string, orderDetails entities.Order) (entities.Order, models.EcomError)
	GetOrderByID(orderId string) (entities.Order, models.EcomError)
	GetAllOrders() ([]entities.Order, models.EcomError)
	GetAllOrderByUserID(ctx context.Context) ([]entities.Order, models.EcomError)
	ConsumeKafkaData(ctx context.Context, data entities.Consume) (entities.Consume, models.EcomError)
}
