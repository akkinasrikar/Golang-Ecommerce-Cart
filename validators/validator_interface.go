package validators

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/gin-gonic/gin"
)

type Validator interface {
	ValidateGetProductReq(context.Context) models.EcomError
	ValidateGetUserDetailsReq(context.Context) models.EcomError
	ValidateCardDetailsReq(ctx *gin.Context) (models.CardDetails, models.EcomError)
	ValidateGetCardDetailsReq(ctx *gin.Context) models.EcomError
	ValidateAddAddressReq(ctx *gin.Context) (models.Address, models.EcomError)
	ValidateGetAddressReq(ctx *gin.Context) models.EcomError
}

type validator struct{}

func NewValidator() Validator {
	return &validator{}
}
