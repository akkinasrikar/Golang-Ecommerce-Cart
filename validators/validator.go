package validators

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func (v *validator) ValidateGetProductReq(ecomCtx context.Context) models.EcomError {	
	return models.EcomError{}
}

func (v *validator) ValidateGetUserDetailsReq(ecomCtx context.Context) models.EcomError {
	return models.EcomError{}
}

func (v *validator) ValidateCardDetailsReq(ctx *gin.Context) (req models.CardDetails, err models.EcomError) {
	err = utils.ValidateUnkownParams(ctx, req)
	if err.Message != nil {
		return models.CardDetails{}, err
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return models.CardDetails{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}

	rules := govalidator.MapData{
		"card_number": []string{"required", "numeric", "len:16"},
		"cvv":         []string{"required", "numeric", "len:3"},
		"name":        []string{"regex:" + constants.RegularExpression.NAME},
		"card_type":   []string{"regex:" + constants.RegularExpression.CardType},
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: rules,
	}

	validator := govalidator.New(opts)
	vErrs := validator.ValidateStruct()
	if len(vErrs) > 0 {
		ecomErr := helper.GetValidationEcomError(vErrs)
		return models.CardDetails{}, ecomErr
	}

	isValidExpiryDate := utils.ValidateCardExpiryDate(req.ExpiryDate)
	if !isValidExpiryDate {
		return models.CardDetails{}, *helper.ErrorParamMissingOrInvalid("Invalid expiry date", "expiry_date")
	}

	return req, models.EcomError{}
}

func (v *validator) ValidateGetCardDetailsReq(ctx *gin.Context) models.EcomError {
	return models.EcomError{}
}

func (v *validator) ValidateAddAddressReq(ctx *gin.Context) (req models.Address, err models.EcomError) {
	err = utils.ValidateUnkownParams(ctx, req)
	if err.Message != nil {
		return models.Address{}, err
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return models.Address{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}

	rules := govalidator.MapData{
		"pincode":  []string{"required", "regex:" + constants.RegularExpression.Pincode},
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: rules,
	}

	validator := govalidator.New(opts)
	vErrs := validator.ValidateStruct()
	if len(vErrs) > 0 {
		ecomErr := helper.GetValidationEcomError(vErrs)
		return models.Address{}, ecomErr
	}

	return req, models.EcomError{}
}

func (v *validator) ValidateGetAddressReq(ctx *gin.Context) models.EcomError {
	return models.EcomError{}
}
