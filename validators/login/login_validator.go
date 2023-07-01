package login

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
)

func (v *loginValidator) ValidateSignUp(ctx *gin.Context) (reqBody entities.SignUp, ecomErr models.EcomError) {
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return entities.SignUp{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}
	return reqBody, models.EcomError{}
}
