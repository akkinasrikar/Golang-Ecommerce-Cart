package validators

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/akkinasrikar/ecommerce-cart/config"
	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func (v *validator) ValidateGetProductReq(ecomCtx context.Context) models.EcomError {
	return models.EcomError{}
}

func (v *validator) ValidateGetOrdersByUserIdReq(ctx *gin.Context) models.EcomError {
	return models.EcomError{}
}

func (v *validator) ValidateGetProductByIdReq(ctx *gin.Context) (int, models.EcomError) {
	id := ctx.Query("id")
	if id == "" {
		return 0, *helper.ErrorParamMissingOrInvalid("Invalid id", "id")
	}
	Id, err := strconv.Atoi(id)
	if err != nil {
		return 0, *helper.ErrorParamMissingOrInvalid("Invalid id", "id")
	}
	return Id, models.EcomError{}
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
		"pincode": []string{"required", "regex:" + constants.RegularExpression.Pincode},
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

func (v *validator) ValidateAddToCartReq(ctx *gin.Context) (req models.AddToCart, err models.EcomError) {
	err = utils.ValidateUnkownParams(ctx, req)
	if err.Message != nil {
		return models.AddToCart{}, err
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return models.AddToCart{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}

	rules := govalidator.MapData{
		"product_id": []string{"required", "regex:" + constants.RegularExpression.ProductID},
		"action":     []string{"required", "regex:" + constants.RegularExpression.Action},
	}

	opts := govalidator.Options{
		Data:  &req,
		Rules: rules,
	}

	validator := govalidator.New(opts)
	vErrs := validator.ValidateStruct()
	if len(vErrs) > 0 {
		ecomErr := helper.GetValidationEcomError(vErrs)
		return models.AddToCart{}, ecomErr
	}

	return req, models.EcomError{}
}

func (v *validator) ValidateGetProductsFromCartReq(ctx *gin.Context) models.EcomError {
	return models.EcomError{}
}

func (v *validator) ValidateOrderProductsReq(ctx *gin.Context) (placeOrder models.PlaceOrder, err models.EcomError) {
	var req models.EcomOrders
	ecomErr := utils.ValidateUnkownParams(ctx, req)
	if ecomErr.Message != nil {
		return models.PlaceOrder{}, ecomErr
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return models.PlaceOrder{}, *helper.ErrorParamMissingOrInvalid("Invalid request body", "body")
	}

	ecomGinCtx, _ := ctx.Get("EcomCtx")
	ecomCtx := ecomGinCtx.(context.Context)
	userDetails, ecomErr := v.Store.GetUserDetails(ecomCtx)
	if ecomErr.Message != nil {
		return models.PlaceOrder{}, ecomErr
	}

	var cartItems entities.ItemsInCart
	unmarshallErr := json.Unmarshal([]byte(userDetails.CartItems), &cartItems)
	if unmarshallErr != nil {
		return models.PlaceOrder{}, *helper.ErrorParamMissingOrInvalid("Invalid cart items", "cart_items")
	}

	if len(cartItems.ItemsID) == 0 {
		return models.PlaceOrder{}, *helper.ErrorParamMissingOrInvalid("No items in cart", "cart_items")
	}

	addressDetails, ecomErr := v.Store.GetAddressById(req.AddressID)
	if ecomErr.Message != nil {
		return models.PlaceOrder{}, ecomErr
	}
	if addressDetails.AddressID == "" {
		return models.PlaceOrder{}, *helper.ErrorParamMissingOrInvalid("Invalid address id", "address_id")
	}

	cardDetails, ecomErr := v.Store.GetCardDetailsById(req.CardId)
	if ecomErr.Message != nil {
		return models.PlaceOrder{}, ecomErr
	}
	if cardDetails.CardID == "" {
		return models.PlaceOrder{}, *helper.ErrorParamMissingOrInvalid("Invalid card id", "card_id")
	}

	decryptCardDetails, decryptErr := utils.DecryptData([]byte(cardDetails.EncryptedData), config.FakeStore.PrivateKey)
	if decryptErr != nil {
		return models.PlaceOrder{}, *helper.ErrorInternalSystemError(decryptErr.Error())
	}

	var cardDetailsJson models.CardDetails
	unmarshallErr = json.Unmarshal(decryptCardDetails, &cardDetailsJson)
	if unmarshallErr != nil {
		return models.PlaceOrder{}, *helper.ErrorInternalSystemError(err.Error())
	}

	placeOrder.AddressID = addressDetails.AddressID
	placeOrder.Address = utils.GetDeliveryAddress(addressDetails)
	placeOrder.CardId = cardDetails.CardID
	placeOrder.CardNumber = cardDetailsJson.CardNumber
	placeOrder.EcomId = userDetails.EcomID
	placeOrder.UsersID = userDetails.UsersID

	return placeOrder, models.EcomError{}
}
