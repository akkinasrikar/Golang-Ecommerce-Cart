package helper

import (
	"errors"

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
