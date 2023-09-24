package validators

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	"github.com/gin-gonic/gin"
)

type Validator interface {
	ValidateGetProductReq(context.Context) models.EcomError
	ValidateGetProductByIdReq(ctx *gin.Context) (int, models.EcomError)
	ValidateGetUserDetailsReq(context.Context) models.EcomError
	ValidateCardDetailsReq(ctx *gin.Context) (models.CardDetails, models.EcomError)
	ValidateGetCardDetailsReq(ctx *gin.Context) models.EcomError
	ValidateAddAddressReq(ctx *gin.Context) (models.Address, models.EcomError)
	ValidateGetAddressReq(ctx *gin.Context) models.EcomError
	ValidateAddToCartReq(ctx *gin.Context) (models.AddToCart, models.EcomError)
	ValidateGetProductsFromCartReq(ctx *gin.Context) models.EcomError
	ValidateOrderProductsReq(ctx *gin.Context) (models.PlaceOrder, models.EcomError)
	ValidateGetOrdersByUserIdReq(ctx *gin.Context) models.EcomError
}

type validator struct {
	Store repositories.RepositoryInterface
}

func NewValidator(store repositories.RepositoryInterface) Validator {
	return &validator{
		Store: store,
	}
}
