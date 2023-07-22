package login

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/gin-gonic/gin"
)

type loginValidator struct{}

func NewValidator() LoginValidator {
	return &loginValidator{}
}

//go:generate mockgen -package mocks -source=login_interface.go -destination=mocks/login_interface_mocks.go
type LoginValidator interface {
	ValidateSignUp(ctx *gin.Context) (models.SignUp, models.EcomError)
	ValidateLogin(ctx *gin.Context) (models.Login, models.EcomError)
}
