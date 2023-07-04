package controllers

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	services "github.com/akkinasrikar/ecommerce-cart/services/login"
	validator "github.com/akkinasrikar/ecommerce-cart/validators/login"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginService   services.LoginService
	loginValidator validator.LoginValidator
	repoService    *gorm.DB
}

func NewLoginHandler(loginService services.LoginService, loginValidator validator.LoginValidator, db *gorm.DB) *LoginHandler {
	return &LoginHandler{
		loginService:   loginService,
		loginValidator: loginValidator,
		repoService:    db,
	}
}

func (lh *LoginHandler) SignUp(ctx *gin.Context) {
	req, ecomError := lh.loginValidator.ValidateSignUp(ctx)
	if ecomError.Message != nil {
		ctx.Error(errors.Wrap(&ecomError, "Error validating request body for Login/SignUp"))
		return
	}
	resp, err := lh.loginService.SignUp(req)
	if err != nil {
		ctx.Error(errors.Wrap(err, "Error in Login/SignUp"))
		return
	}
	ctx.JSON(200, resp)
}

func (lh *LoginHandler) Login(ctx *gin.Context) {
	req, ecomError := lh.loginValidator.ValidateLogin(ctx)
	if ecomError.Message != nil {
		ctx.Error(errors.Wrap(&ecomError, "Error validating request body for Login"))
		return
	}
	resp, err := lh.loginService.Login(req)
	if err != nil {
		ctx.Error(errors.Wrap(err, "Error in Login/Login"))
		return
	}
	ctx.JSON(200, resp)
}

func (lh *LoginHandler) HomePage(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to Ecommerce Cart",
	})
}
