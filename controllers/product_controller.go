package controllers

import (
	"context"
	"net/http"

	"github.com/akkinasrikar/ecommerce-cart/repositories"
	services "github.com/akkinasrikar/ecommerce-cart/services/products"
	validator "github.com/akkinasrikar/ecommerce-cart/validators"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productValidator validator.Validator
	ecomService      services.Products
	store            repositories.RepositoryInterface
}

func NewProductHandler(productValidator validator.Validator, ecomService services.Products) *ProductHandler {
	return &ProductHandler{
		productValidator: productValidator,
		ecomService:      ecomService,
	}
}

func (ph *ProductHandler) GetProducts(ctx *gin.Context) {
	ecomError := ph.productValidator.ValidateGetProductReq(ctx)
	if ecomError.Message != nil {
		ecomError := helper.SetInternalError(ecomError.Message.Error())
		ctx.JSON(int(ecomError.ErrorType.Code), &ecomError)
		return
	}
	resp, err := ph.ecomService.GetProducts(ctx)
	if err.Message != nil {
		ecomErr := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(ecomErr.ErrorType.Code), &ecomErr)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) GetUserDetails(ctx *gin.Context) {
	ecomError := ph.productValidator.ValidateGetUserDetailsReq(ctx)
	if ecomError.Message != nil {
		ecomError := helper.SetInternalError(ecomError.Message.Error())
		ctx.JSON(int(ecomError.ErrorType.Code), &ecomError)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.GetUserDetails(ecomCtx)
	if err.Message != nil {
		ecomErr := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(ecomErr.ErrorType.Code), &ecomErr)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) CardDetails(ctx *gin.Context) {
	req, ecomError := ph.productValidator.ValidateCardDetailsReq(ctx)
	if ecomError.Message != nil {
		ecomError := helper.SetInternalError(ecomError.Message.Error())
		ctx.JSON(int(ecomError.ErrorType.Code), &ecomError)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.CardDetails(ecomCtx, req)
	if err.Message != nil {
		ecomErr := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(ecomErr.ErrorType.Code), &ecomErr)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) GetCardDetails(ctx *gin.Context) {
	ecomError := ph.productValidator.ValidateGetCardDetailsReq(ctx)
	if ecomError.Message != nil {
		ecomError := helper.SetInternalError(ecomError.Message.Error())
		ctx.JSON(int(ecomError.ErrorType.Code), &ecomError)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.GetCardDetails(ecomCtx)
	if err.Message != nil {
		ecomErr := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(ecomErr.ErrorType.Code), &ecomErr)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
