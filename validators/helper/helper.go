package helper

import (
	"errors"
	"net/url"

	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models"
)

var ErrorParamMissingOrInvalid = func(msg string, param string) *models.EcomError {
	return &models.EcomError{
		Code:    422,
		Message: errors.New(msg),
		Type:    "param_missing_or_invalid",
		Param:   param,
	}
}

var ErrorInternalSystemError = func(msg string) *models.EcomError {
	return &models.EcomError{
		Code:    int64(constants.ErrorType.INTERNAL_SYSTEM_ERROR.Code),
		Message: errors.New(msg),
		Type:    constants.ErrorType.INTERNAL_SYSTEM_ERROR.Name,
	}
}

var ErrorDownStreamError = func() *models.EcomError {
	return &models.EcomError{
		Code:    int64(constants.ErrorType.DOWNSTREAM_ERROR.Code),
		Message: errors.New("downstream error"),
		Type:    constants.ErrorType.DOWNSTREAM_ERROR.Name,
	}
}

func SetInternalError(errMsg string) models.EcomErrorResponse {
	zwErrBody := models.EcomErrorBody{
		Message: errMsg,
		Type:    constants.ErrorType.INTERNAL_SYSTEM_ERROR.Name,
		Code:    int64(constants.ErrorType.INTERNAL_SYSTEM_ERROR.Code),
	}
	return models.EcomErrorResponse{
		ErrorType: zwErrBody,
	}
}

func GetValidationEcomError(e url.Values) models.EcomError {
	var ecomErr models.EcomError
	for key, value := range e {
		ecomErr = *ErrorParamMissingOrInvalid(value[0], key)
		break
	}
	return ecomErr
}

func ErrorUnknownParam(param string) *models.EcomError {
	return &models.EcomError{
		Code:    422,
		Message: errors.New("unknown param"),
		Type:    "unknown_param",
		Param:   param,
	}
}
