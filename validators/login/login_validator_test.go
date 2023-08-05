package login

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/gin-gonic/gin"
)

func Test_loginValidator_ValidateSignUp(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	signUpReq := models.SignUp{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "test12345678",
	}
	tests := []struct {
		name        string
		v           *loginValidator
		args        args
		wantReqBody models.SignUp
		wantEcomErr models.EcomError
	}{
		{
			name: "Success",
			v:    &loginValidator{},
			args: args{
				ctx: utils.SetContext(),
			},
			wantReqBody: signUpReq,
			wantEcomErr: models.EcomError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBodyCtx, _ := json.Marshal(signUpReq)
			tt.args.ctx.Request = &http.Request{
				Header: make(http.Header),
				Body:   ioutil.NopCloser(bytes.NewBuffer(reqBodyCtx)),
			}
			v := &loginValidator{}
			gotReqBody, gotEcomErr := v.ValidateSignUp(tt.args.ctx)
			if !reflect.DeepEqual(gotReqBody, tt.wantReqBody) {
				t.Errorf("loginValidator.ValidateSignUp() gotReqBody = %v, want %v", gotReqBody, tt.wantReqBody)
			}
			if gotEcomErr.Message != nil && (gotEcomErr.Message != tt.wantEcomErr.Message) {
				t.Errorf("loginValidator.ValidateSignUp() gotEcomErr = %v, want %v", gotEcomErr, tt.wantEcomErr)
			}
		})
	}
}

func Test_loginValidator_ValidateLogin(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	loginReq := models.Login{
		Name:     "test",
		Password: "test12345678",
	}
	tests := []struct {
		name        string
		v           *loginValidator
		args        args
		wantReqBody models.Login
		wantEcomErr models.EcomError
	}{
		{
			name: "Success",
			v:    &loginValidator{},
			args: args{
				ctx: utils.SetContext(),
			},
			wantReqBody: loginReq,
			wantEcomErr: models.EcomError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBodyCtx, _ := json.Marshal(loginReq)
			tt.args.ctx.Request = &http.Request{
				Header: make(http.Header),
				Body:   ioutil.NopCloser(bytes.NewBuffer(reqBodyCtx)),
			}
			v := &loginValidator{}
			gotReqBody, gotEcomErr := v.ValidateLogin(tt.args.ctx)
			if !reflect.DeepEqual(gotReqBody, tt.wantReqBody) {
				t.Errorf("loginValidator.ValidateLogin() gotReqBody = %v, want %v", gotReqBody, tt.wantReqBody)
			}
			if gotEcomErr.Message != nil && (gotEcomErr.Message != tt.wantEcomErr.Message) {
				t.Errorf("loginValidator.ValidateLogin() gotEcomErr = %v, want %v", gotEcomErr, tt.wantEcomErr)
			}
		})
	}
}
