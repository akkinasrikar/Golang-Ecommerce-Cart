package login

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/gin-gonic/gin"
)

type loginValidator struct {}

func NewValidator() LoginValidator {
	return &loginValidator{}
}

type LoginValidator interface {
	ValidateSignUp(ctx *gin.Context) (entities.SignUp, models.EcomError)
	ValidateLogin(ctx *gin.Context) (entities.Login, models.EcomError)
}
