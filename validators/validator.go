package validators

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
)

func (v *validator) ValidateGetProductReq(ecomCtx context.Context) models.EcomError {
	return models.EcomError{}
}

func (v *validator) ValidateGetUserDetailsReq(ecomCtx context.Context) models.EcomError {
	return models.EcomError{}
}


func (v *validator) ValidateCardDetailsReq(ctx *gin.Context) (req models.CardDetails, err models.EcomError) {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return models.CardDetails{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}
	return req, models.EcomError{}
}
