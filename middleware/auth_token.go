package middleware

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/gin-gonic/gin"
)

func ValidateJwtAuthToken() func(*gin.Context) {
	return func(ctx *gin.Context) {
		var err error
		var header models.AuthData
		ctx.ShouldBindHeader(&header)
	}
}

func SetRequestContext(authData models.AuthData) context.Context {
	var ecomCtx context.Context
	ecomCtx = context.Background()
	ecomCtx = context.WithValue(ecomCtx, models.EcomctxKey("AuthData"), authData)
	return ecomCtx
}
