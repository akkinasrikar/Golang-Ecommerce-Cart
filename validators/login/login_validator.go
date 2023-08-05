package login

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func (v *loginValidator) ValidateSignUp(ctx *gin.Context) (reqBody models.SignUp, ecomErr models.EcomError) {
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return models.SignUp{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}

	rules := getRulesForSignUp()

	opts := govalidator.Options{
		Data:  &reqBody,
		Rules: rules,
		Messages: govalidator.MapData{
			"user_name":  []string{"required:Name is required", "between:Name should be between 3 to 20 characters"},
			"user_email": []string{"required:Email is required", "email:Email should be valid"},
			"password":   []string{"required:Password is required", "min:Password should be minimum 8 characters", "max:Password should be maximum 20 characters"},
		},
	}

	vErrs := govalidator.New(opts).ValidateStruct()
	if len(vErrs) > 0 {
		ecomErr := helper.GetValidationEcomError(vErrs)
		return models.SignUp{}, ecomErr
	}

	return reqBody, models.EcomError{}
}

func getRulesForSignUp() govalidator.MapData {
	return govalidator.MapData{
		"user_name":  []string{"required", "between:3,20"},
		"user_email": []string{"required", "email"},
		"password":   []string{"required", "min:8", "max:20"},
	}
}

func (v *loginValidator) ValidateLogin(ctx *gin.Context) (reqBody models.Login, ecomErr models.EcomError) {
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		return models.Login{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}
	return reqBody, models.EcomError{}
}
