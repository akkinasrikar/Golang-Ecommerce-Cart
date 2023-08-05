package validators

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/models"
)

type Validator interface {
	ValidateGetProductReq(context.Context) models.EcomError
}

type validator struct{}

func NewValidator() Validator {
	return &validator{}
}
