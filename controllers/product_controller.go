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

// SeedData
func (ph *ProductHandler) SeedData(ctx *gin.Context) {
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	err := ph.ecomService.SeedData(ecomCtx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
	}
	ctx.JSON(http.StatusOK, "Data seeded successfully")
}

func (ph *ProductHandler) GetProducts(ctx *gin.Context) {
	err := ph.productValidator.ValidateGetProductReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	resp, err := ph.ecomService.GetProducts(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) GetProductById(ctx *gin.Context) {
	id, err := ph.productValidator.ValidateGetProductByIdReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	resp, err := ph.ecomService.GetProductById(ctx,id)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ctx.Header("Content-Type", "text/html")
	ctx.String(http.StatusOK, resp)
}

func (ph *ProductHandler) GetUserDetails(ctx *gin.Context) {
	err := ph.productValidator.ValidateGetUserDetailsReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
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
	req, err := ph.productValidator.ValidateCardDetailsReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.CardDetails(ecomCtx, req)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) GetCardDetails(ctx *gin.Context) {
	err := ph.productValidator.ValidateGetCardDetailsReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.GetCardDetails(ecomCtx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
 
func (ph *ProductHandler) AddAddress(ctx *gin.Context) {
	req, err := ph.productValidator.ValidateAddAddressReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.AddAddress(ecomCtx, req)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) GetAddress(ctx *gin.Context) {
	err := ph.productValidator.ValidateGetAddressReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.GetAddress(ecomCtx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) AddOrDeleteToCart(ctx *gin.Context) {
	req, err := ph.productValidator.ValidateAddToCartReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.AddOrDeleteToCart(ecomCtx, req)
	if err.Message != nil {
		ctx.JSON(int(err.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) GetProductsFromCart(ctx *gin.Context) {
	err := ph.productValidator.ValidateGetProductsFromCartReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.GetProductsFromCart(ecomCtx)
	if err.Message != nil {
		ctx.JSON(int(err.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) OrderProducts(ctx *gin.Context) {
	req, err := ph.productValidator.ValidateOrderProductsReq(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	resp, err := ph.ecomService.OrderProducts(ecomCtx, req)
	if err.Message != nil {
		ctx.JSON(int(err.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
