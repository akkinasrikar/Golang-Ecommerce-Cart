package controllers

import (
	"github.com/pkg/errors"

	services "github.com/akkinasrikar/ecommerce-cart/services/login"
	validator "github.com/akkinasrikar/ecommerce-cart/validators/login"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginService   services.LoginService
	loginValidator validator.LoginValidator
}

func NewLoginHandler(loginService services.LoginService, loginValidator validator.LoginValidator) *LoginHandler {
	return &LoginHandler{
		loginService:   loginService,
		loginValidator: loginValidator,
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
