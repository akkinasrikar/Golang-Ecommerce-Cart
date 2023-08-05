package services

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
)

func (p *products) GetProducts(ctx context.Context) (responses.ItemsResponse, models.EcomError) {
	var itemsResponse responses.ItemsResponse
	resp, err := p.APIProvider.GetItems(ctx)
	if err.Message != nil {
		return itemsResponse, err
	}
	for _, item := range resp {
		itemsResponse = append(itemsResponse, responses.Items{
			Id:          item.Id,
			Title:       item.Title,
			Price:       item.Price,
			Description: item.Description,
			Category:    item.Category,
			Image:       item.Image,
			Rating: responses.Rating{
				Rate:  item.Rating.Rate,
				Count: item.Rating.Count,
			},
		})
	}
	return itemsResponse, models.EcomError{}
}
