package validators

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/models"
)

func (v *validator) ValidateGetProductReq(ecomCtx context.Context) models.EcomError {
	return models.EcomError{}
}

func (v *validator) ValidateGetUserDetailsReq(ecomCtx context.Context) models.EcomError {	
	return models.EcomError{}
}
