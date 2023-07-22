package controllers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
	services "github.com/akkinasrikar/ecommerce-cart/services/login"
	smocks "github.com/akkinasrikar/ecommerce-cart/services/login/mocks"
	validator "github.com/akkinasrikar/ecommerce-cart/validators/login"
	vmocks "github.com/akkinasrikar/ecommerce-cart/validators/login/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestLoginHandler_SignUp(t *testing.T) {
	type fields struct {
		loginService   services.LoginService
		loginValidator validator.LoginValidator
	}
	type args struct {
		ctx *gin.Context
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Set("EcomCtx", context.Background())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		fields        fields
		args          args
		mockService   func(ctrl *gomock.Controller) *smocks.MockLoginService
		mockValidator func(ctrl *gomock.Controller) *vmocks.MockLoginValidator
	}{
		{
			name: "Success",
			mockService: func(ctrl *gomock.Controller) *smocks.MockLoginService {
				mockService := smocks.NewMockLoginService(ctrl)
				mockService.EXPECT().SignUp(gomock.Any()).Return(responses.SingUp{}, models.EcomError{})
				return mockService
			},
			mockValidator: func(ctrl *gomock.Controller) *vmocks.MockLoginValidator {
				mockValidator := vmocks.NewMockLoginValidator(ctrl)
				mockValidator.EXPECT().ValidateSignUp(gomock.Any()).Return(models.SignUp{}, models.EcomError{})
				return mockValidator
			},
			args: args{
				ctx: ctx,
			},
		},
		{
			name: "Validation Error",
			mockService: func(ctrl *gomock.Controller) *smocks.MockLoginService {
				mockService := smocks.NewMockLoginService(ctrl)
				return mockService
			},
			mockValidator: func(ctrl *gomock.Controller) *vmocks.MockLoginValidator {
				mockValidator := vmocks.NewMockLoginValidator(ctrl)
				mockValidator.EXPECT().ValidateSignUp(gomock.Any()).Return(models.SignUp{}, models.EcomError{
					Message: fmt.Errorf("Validation Error"),
				})
				return mockValidator
			},
			args: args{
				ctx: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lh := &LoginHandler{
				loginService:   tt.mockService(ctrl),
				loginValidator: tt.mockValidator(ctrl),
			}
			lh.SignUp(tt.args.ctx)
		})
	}
}

func TestLoginHandler_Login(t *testing.T) {
	type fields struct {
		loginService   services.LoginService
		loginValidator validator.LoginValidator
	}
	type args struct {
		ctx *gin.Context
	}

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Set("EcomCtx", context.Background())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		fields        fields
		args          args
		mockService   func(ctrl *gomock.Controller) *smocks.MockLoginService
		mockValidator func(ctrl *gomock.Controller) *vmocks.MockLoginValidator
	}{
		{
			name: "Success",
			mockService: func(ctrl *gomock.Controller) *smocks.MockLoginService {
				mockService := smocks.NewMockLoginService(ctrl)
				mockService.EXPECT().Login(gomock.Any()).Return(responses.Login{}, models.EcomError{})
				return mockService
			},
			mockValidator: func(ctrl *gomock.Controller) *vmocks.MockLoginValidator {
				mockValidator := vmocks.NewMockLoginValidator(ctrl)
				mockValidator.EXPECT().ValidateLogin(gomock.Any()).Return(models.Login{}, models.EcomError{})
				return mockValidator
			},
			args: args{
				ctx: ctx,
			},
		},
		{
			name: "Validation Error",
			mockService: func(ctrl *gomock.Controller) *smocks.MockLoginService {
				mockService := smocks.NewMockLoginService(ctrl)
				return mockService
			},
			mockValidator: func(ctrl *gomock.Controller) *vmocks.MockLoginValidator {
				mockValidator := vmocks.NewMockLoginValidator(ctrl)
				mockValidator.EXPECT().ValidateLogin(gomock.Any()).Return(models.Login{}, models.EcomError{
					Message: fmt.Errorf("Validation Error"),
				})
				return mockValidator
			},
			args: args{
				ctx: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lh := &LoginHandler{
				loginService:   tt.mockService(ctrl),
				loginValidator: tt.mockValidator(ctrl),
			}
			lh.Login(tt.args.ctx)
		})
	}
}
