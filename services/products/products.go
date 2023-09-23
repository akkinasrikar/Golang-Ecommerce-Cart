package services

import (
	"context"
	"encoding/json"

	"github.com/akkinasrikar/ecommerce-cart/config"
	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
)

func (p *products) SeedData(ctx context.Context) models.EcomError {
	item, err := p.Store.GetProductById(1)
	if item.ItemID != 0 || err.Message != nil {
		return *helper.ErrorInternalSystemError("data already seeded")
	}
	apiResponse, err := p.APIProvider.GetItems(ctx)
	if err.Message != nil {
		return *helper.ErrorInternalSystemError(err.Message.Error())
	}
	for _, value := range apiResponse {
		item := entities.Item{
			ItemID:          value.Id,
			ItemTitle:       value.Title,
			ItemPrice:       value.Price,
			ItemImage:       value.Image,
			ItemRating:      value.Rating.Rate,
			ItemDescription: value.Description,
			ItemCategory:    value.Category,
			ItemCount:       value.Rating.Count,
		}
		_, err := p.Store.CreateProduct(item)
		if err.Message != nil {
			return *helper.ErrorInternalSystemError(err.Message.Error())
		}
		asynqErr := p.ProducAsynqService.ProductImageResize(ctx, value.Id)
		if asynqErr != nil {
			return *helper.ErrorInternalSystemError(asynqErr.Error())
		}
	}
	return models.EcomError{}
}

func (p *products) GetProducts(ctx context.Context) ([]entities.Item, models.EcomError) {
	var items []entities.Item
	items, err := p.Store.GetAllProducts()
	if err.Message != nil {
		return items, err
	}
	return items, models.EcomError{}
}

func (p *products) GetProductById(ctx context.Context, id int) (string, models.EcomError) {
	item, err := p.Store.GetProductById(id)
	if err.Message != nil {
		return "", err
	}
	htmlResponse := utils.GenerateHtmlResponse(item.ImageBase64, item)
	return htmlResponse, models.EcomError{}
}

func (p *products) GetUserDetails(ctx context.Context) (models.EcomUsers, models.EcomError) {
	user, err := p.Store.GetUserDetails(ctx)
	if err.Message != nil {
		return models.EcomUsers{}, err
	}
	cartItems, cartErr := utils.UnmarshallCartItems(user.CartItems)
	if cartErr != nil {
		return models.EcomUsers{}, *helper.ErrorInternalSystemError(err.Error())
	}
	userDetails := models.EcomUsers{
		EcomID:       user.EcomID,
		AccountName:  user.AccountName,
		WalletAmount: user.WalletAmount,
		UsersID:      user.UsersID,
		CartItems:    cartItems.ItemsID,
	}
	return userDetails, models.EcomError{}
}

func (p *products) CardDetails(ctx context.Context, req models.CardDetails) (models.CardDetails, models.EcomError) {
	userDetails, ecomErr := p.Store.GetUserDetails(ctx)
	if ecomErr.Message != nil {
		return req, ecomErr
	}

	cardId := utils.GenerateCardId()
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return req, *helper.ErrorInternalSystemError(err.Error())
	}

	encyptedData, err := utils.EncryptData(jsonReq, config.FakeStore.PublicKey)
	if err != nil {
		return req, *helper.ErrorInternalSystemError(err.Error())
	}

	encryptedCardDetails := entities.CardDetails{
		CardID:        cardId,
		EncryptedData: encyptedData,
		EcomId:        userDetails.EcomID,
	}

	_, ecomErr = p.Store.CreateCardDetails(encryptedCardDetails)
	if ecomErr.Message != nil {
		return req, ecomErr
	}

	return req, models.EcomError{}
}

func (p *products) GetCardDetails(ctx context.Context) ([]models.CardDetails, models.EcomError) {
	var decryptedCardDetails []models.CardDetails

	userDetails, ecomErr := p.Store.GetUserDetails(ctx)
	if ecomErr.Message != nil {
		return []models.CardDetails{}, ecomErr
	}

	var cardDetails []entities.CardDetails
	cardDetails, err := p.Store.GetCardDetails(userDetails)
	if err.Message != nil {
		return []models.CardDetails{}, err
	}

	for _, cardDetail := range cardDetails {
		decryptedData, err := utils.DecryptData([]byte(cardDetail.EncryptedData), config.FakeStore.PrivateKey)
		if err != nil {
			return []models.CardDetails{}, *helper.ErrorInternalSystemError(err.Error())
		}
		var decryptedCardDetail models.CardDetails
		err = json.Unmarshal(decryptedData, &decryptedCardDetail)
		if err != nil {
			return []models.CardDetails{}, *helper.ErrorInternalSystemError(err.Error())
		}
		decryptedCardDetail.CardID = cardDetail.CardID
		decryptedCardDetails = append(decryptedCardDetails, decryptedCardDetail)
	}
	return decryptedCardDetails, models.EcomError{}
}

func (p *products) AddAddress(ctx context.Context, req models.Address) (entities.DeliveryAddress, models.EcomError) {
	userDetails, ecomErr := p.Store.GetUserDetails(ctx)
	if ecomErr.Message != nil {
		return entities.DeliveryAddress{}, ecomErr
	}

	addressObject := entities.DeliveryAddress{
		EcomID:    userDetails.EcomID,
		AddressID: utils.GenerateAddressId(),
		HouseNo:   req.HouseNo,
		Street:    req.Street,
		City:      req.City,
		State:     req.State,
		Pincode:   req.Pincode,
	}

	address, ecomErr := p.Store.CreateAddress(addressObject)
	if ecomErr.Message != nil {
		return entities.DeliveryAddress{}, ecomErr
	}

	return address, models.EcomError{}
}

func (p *products) GetAddress(ctx context.Context) ([]entities.DeliveryAddress, models.EcomError) {
	var addresses []entities.DeliveryAddress

	userDetails, ecomErr := p.Store.GetUserDetails(ctx)
	if ecomErr.Message != nil {
		return []entities.DeliveryAddress{}, ecomErr
	}

	var addressDetails []entities.DeliveryAddress
	addressDetails, err := p.Store.GetAddress(userDetails)
	if err.Message != nil {
		return []entities.DeliveryAddress{}, err
	}

	for _, addressDetail := range addressDetails {
		address := entities.DeliveryAddress{
			AddressID: addressDetail.AddressID,
			HouseNo:   addressDetail.HouseNo,
			Street:    addressDetail.Street,
			City:      addressDetail.City,
			State:     addressDetail.State,
			Pincode:   addressDetail.Pincode,
			EcomID:    addressDetail.EcomID,
		}
		addresses = append(addresses, address)
	}
	return addresses, models.EcomError{}
}

func (p *products) AddOrDeleteToCart(ctx context.Context, req models.AddToCart) (models.CartResponse, models.EcomError) {
	var cartItems entities.ItemsInCart
	var cartResponse models.CartResponse
	userDetails, ecomErr := p.Store.GetUserDetails(ctx)
	if ecomErr.Message != nil {
		return models.CartResponse{}, ecomErr
	}
	err := json.Unmarshal([]byte(userDetails.CartItems), &cartItems)
	if err != nil {
		return models.CartResponse{}, *helper.ErrorInternalSystemError(err.Error())
	}
	if req.Action == constants.ProductConstants.ADDITION {
		item, ecomErr := p.Store.AddToCart(userDetails, req.ProductId)
		if ecomErr.Message != nil {
			return models.CartResponse{}, ecomErr
		}
		if item.ItemID == 0 {
			return models.CartResponse{}, *helper.ErrorParamMissingOrInvalid("invalid product id", "product_id")
		}
		cartItems.ItemsID = append(cartItems.ItemsID, item.ItemID)
		cartItemsJson, err := json.Marshal(cartItems)
		if err != nil {
			return models.CartResponse{}, *helper.ErrorInternalSystemError(err.Error())
		}
		userDetails.CartItems = string(cartItemsJson)
		_, ecomErr = p.Store.UpdateEcomAccount(userDetails, userDetails.EcomID)
		if ecomErr.Message != nil {
			return models.CartResponse{}, ecomErr
		}
		cartResponse.Action = constants.ProductConstants.ADDITION
		cartResponse.ProductID = req.ProductId
		cartResponse.Message = "Successfully added to cart!"
	} else if req.Action == constants.ProductConstants.DELETION {
		isValidProductId := false
		for i, v := range cartItems.ItemsID {
			if v == req.ProductId {
				cartItems.ItemsID = append(cartItems.ItemsID[:i], cartItems.ItemsID[i+1:]...)
				isValidProductId = true
				break
			}
		}
		if !isValidProductId {
			return models.CartResponse{}, *helper.ErrorParamMissingOrInvalid("invalid product id", "product_id")
		}
		cartItemsJson, err := json.Marshal(cartItems)
		if err != nil {
			return models.CartResponse{}, *helper.ErrorInternalSystemError(err.Error())
		}
		userDetails.CartItems = string(cartItemsJson)
		_, ecomErr = p.Store.UpdateEcomAccount(userDetails, userDetails.EcomID)
		if ecomErr.Message != nil {
			return models.CartResponse{}, ecomErr
		}
		cartResponse.Action = constants.ProductConstants.DELETION
		cartResponse.ProductID = req.ProductId
		cartResponse.Message = "Successfully deleted from cart!"
	}
	return cartResponse, models.EcomError{}
}

func (p *products) GetProductsFromCart(ctx context.Context) ([]entities.Item, models.EcomError) {
	var items []entities.Item
	var cartItems entities.ItemsInCart
	userDetails, ecomErr := p.Store.GetUserDetails(ctx)
	if ecomErr.Message != nil {
		return []entities.Item{}, ecomErr
	}
	err := json.Unmarshal([]byte(userDetails.CartItems), &cartItems)
	if err != nil {
		return []entities.Item{}, *helper.ErrorInternalSystemError(err.Error())
	}
	for _, v := range cartItems.ItemsID {
		item, ecomErr := p.Store.GetProductFromCart(v)
		if ecomErr.Message != nil {
			return []entities.Item{}, ecomErr
		}
		items = append(items, item)
	}
	return items, models.EcomError{}
}

func (p *products) OrderProducts(ctx context.Context, req models.PlaceOrder) (models.EcomOrderResponse, models.EcomError) {
	var cartItems entities.ItemsInCart
	var orderResponse models.EcomOrderResponse
	var orderDetails []models.OrderDetails

	userDetails, ecomErr := p.Store.GetUserDetails(ctx)
	if ecomErr.Message != nil {
		return models.EcomOrderResponse{}, ecomErr
	}

	err := json.Unmarshal([]byte(userDetails.CartItems), &cartItems)
	if err != nil {
		return models.EcomOrderResponse{}, *helper.ErrorInternalSystemError(err.Error())
	}

	// last four digits of card adn remaining xxxxx
	cardNumber := utils.FormatCardNumber(req.CardNumber)

	for _, value := range cartItems.ItemsID {
		item, ecomErr := p.Store.GetProductFromCart(value)
		if ecomErr.Message != nil {
			return models.EcomOrderResponse{}, ecomErr
		}

		OrderObject := entities.Order{
			OrderID:        utils.GenerateOrderId(),
			OrderStatus:    constants.ProductConstants.SUCCESS,
			OrderAmount:    int64(item.ItemPrice),
			OrderDate:      utils.GenerateCurrentDate(),
			OrderName:      item.ItemTitle,
			PaymentMode:    constants.ProductConstants.CARD,
			DeliveryStatus: constants.ProductConstants.ONTIME,
			DeliveryDate:   utils.GenerateRandomDate(),
			AddressID:      req.AddressID,
			CardID:         req.CardId,
			EcomID:         userDetails.EcomID,
			UsersID:        userDetails.UsersID,
		}

		OrderedObject, ecomErr := p.Store.CreateOrder(OrderObject)
		if ecomErr.Message != nil {
			return models.EcomOrderResponse{}, ecomErr
		}

		orderDetails = append(orderDetails, models.OrderDetails{
			OrderID:      OrderedObject.OrderID,
			Amount:       OrderedObject.OrderAmount,
			ProductName:  OrderedObject.OrderName,
			OrderedDate:  OrderedObject.OrderDate,
			DeliveryDate: OrderedObject.DeliveryDate,
			Address:      req.Address,
			CardNUmber:   cardNumber,
		})
	}
	orderResponse.Orders = orderDetails
	orderResponse.Message = "Successfully ordered!"

	cartItems.ItemsID = []int{}
	cartItemsJson, err := json.Marshal(cartItems)
	if err != nil {
		return models.EcomOrderResponse{}, *helper.ErrorInternalSystemError(err.Error())
	}
	userDetails.CartItems = string(cartItemsJson)
	_, ecomErr = p.Store.UpdateEcomAccount(userDetails, userDetails.EcomID)
	if ecomErr.Message != nil {
		return models.EcomOrderResponse{}, ecomErr
	}
	return orderResponse, models.EcomError{}
}
