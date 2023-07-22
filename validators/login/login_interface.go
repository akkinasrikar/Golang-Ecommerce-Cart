package login

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/gin-gonic/gin"
)

type loginValidator struct{}

func NewValidator() LoginValidator {
	return &loginValidator{}
}

type LoginValidator interface {
	ValidateSignUp(ctx *gin.Context) (models.SignUp, models.EcomError)
	ValidateLogin(ctx *gin.Context) (models.Login, models.EcomError)
}
