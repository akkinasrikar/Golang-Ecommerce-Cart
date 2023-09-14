package services

import (
	"context"
	"encoding/json"

	"github.com/akkinasrikar/ecommerce-cart/config"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
)

func (p *products) GetProducts(ctx context.Context) ([]entities.Item, models.EcomError) {
	var items []entities.Item
	items, err := p.Store.GetAllProducts()
	if err.Message != nil {
		return items, err
	}
	return items, models.EcomError{}
}

func (p *products) GetUserDetails(ctx context.Context) (entities.EcomUsers, models.EcomError) {
	var user entities.EcomUsers
	user, err := p.Store.GetUserDetails(ctx)
	if err.Message != nil {
		return user, err
	}
	return user, models.EcomError{}
}

func (p *products) CardDetails(ctx context.Context, req models.CardDetails) (models.CardDetails, models.EcomError) {
	userDetails, ecomErr := p.Store.GetUserDetails(ctx)
	if ecomErr.Message != nil {
		return req, ecomErr
	}

	req.CardId = utils.GenerateCardId()
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return req, *helper.ErrorInternalSystemError(err.Error())
	}

	encyptedData, err := utils.EncryptData(jsonReq, config.FakeStore.PublicKey)
	if err != nil {
		return req, *helper.ErrorInternalSystemError(err.Error())
	}

	encryptedCardDetails := entities.CardDetails{
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
		decryptedCardDetails = append(decryptedCardDetails, decryptedCardDetail)
	}
	return decryptedCardDetails, models.EcomError{}
}
