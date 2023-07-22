package login

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
)

func (v *loginValidator) ValidateSignUp(ctx *gin.Context) (reqBody models.SignUp, ecomErr models.EcomError) {
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return models.SignUp{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}
	return reqBody, models.EcomError{}
}

func (v *loginValidator) ValidateLogin(ctx *gin.Context) (reqBody models.Login, ecomErr models.EcomError) {
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return models.Login{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}
	return reqBody, models.EcomError{}
}
