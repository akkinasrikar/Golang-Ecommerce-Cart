package controllers

import (
	"context"
	"net/http"

	"gorm.io/gorm"

	"github.com/akkinasrikar/ecommerce-cart/models"
	services "github.com/akkinasrikar/ecommerce-cart/services/login"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	validator "github.com/akkinasrikar/ecommerce-cart/validators/login"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginService   services.LoginService
	loginValidator validator.LoginValidator
}

func NewLoginHandler(loginService services.LoginService, loginValidator validator.LoginValidator, db *gorm.DB) *LoginHandler {
	return &LoginHandler{
		loginService:   loginService,
		loginValidator: loginValidator,
	}
}

func (lh *LoginHandler) SignUp(ctx *gin.Context) {
	req, ecomError := lh.loginValidator.ValidateSignUp(ctx)
	if ecomError.Message != nil {
		ecomError := helper.SetInternalError(ecomError.Message.Error())
		ctx.JSON(int(ecomError.ErrorType.Code), &ecomError)
		return
	}
	resp, err := lh.loginService.SignUp(req)
	if err.Message != nil {
		ecomErr := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(ecomErr.ErrorType.Code), &ecomErr)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (lh *LoginHandler) Login(ctx *gin.Context) {
	req, ecomError := lh.loginValidator.ValidateLogin(ctx)
	if ecomError.Message != nil {
		ecomError := helper.SetInternalError(ecomError.Message.Error())
		ctx.JSON(int(ecomError.ErrorType.Code), &ecomError)
		return
	}
	resp, err := lh.loginService.Login(req)
	if err.Message != nil {
		ecomErr := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(ecomErr.ErrorType.Code), &ecomErr)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (lh *LoginHandler) HomePage(ctx *gin.Context) {
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	authData := ecomCtx.Value(models.EcomctxKey("AuthData")).(models.AuthData)
	ctx.JSON(200, gin.H{
		"message": "Welcome " + authData.UserName + " to HomePage",
	})
}
