package middleware

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateJwtAuthToken() func(*gin.Context) {
	return func(ctx *gin.Context) {
		var header models.AuthData
		ctx.ShouldBindHeader(&header)
		ecomCtx := SetRequestContext(header)

		if header.Authorization == "" {
			ecomErr := setAuthError(constants.ErrorMessage.AUTHENTICATION_REQUIRED)
			ctx.JSON(constants.ErrorType.AUTHENTICATION_ERROR.Code, ecomErr)
			ctx.Abort()
			return
		}
		secret_key := "testing"
		token, err := jwt.Parse(header.Authorization, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret_key), nil
		})
		if err != nil {
			ecomErr := setAuthError(constants.ErrorMessage.INVALID_AUTHENTICATION)
			ctx.JSON(constants.ErrorType.AUTHENTICATION_ERROR.Code, ecomErr)
			ctx.Abort()
			return
		}
		if !token.Valid {
			ecomErr := setAuthError(constants.ErrorMessage.AUTHENTICATION_REQUIRED)
			ctx.JSON(constants.ErrorType.AUTHENTICATION_ERROR.Code, ecomErr)
			ctx.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ecomErr := setAuthError(constants.ErrorMessage.AUTHENTICATION_REQUIRED)
			ctx.JSON(constants.ErrorType.AUTHENTICATION_ERROR.Code, ecomErr)
			ctx.Abort()
			return
		}
		ecomCtx = context.WithValue(ecomCtx, models.EcomctxKey("Claims"), claims)
		ctx.Set("ecomCtx", ecomCtx)
		ctx.Next()
	}
}

func SetRequestContext(authData models.AuthData) context.Context {
	var ecomCtx context.Context
	ecomCtx = context.Background()
	ecomCtx = context.WithValue(ecomCtx, models.EcomctxKey("AuthData"), authData)
	return ecomCtx
}

func setAuthError(errMsg string) models.EcomErrorResponse {
	zwErrBody := models.EcomErrorBody{
		Message: errMsg,
		Type:    constants.ErrorType.AUTHENTICATION_ERROR.Name,
	}
	return models.EcomErrorResponse{
		ErrorType: zwErrBody,
	}
}

func SetAuthError(ctx *gin.Context) {
}
