package services

import (
	"context"

	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
)

func (p *products) GetProducts(ctx context.Context) ([]entities.Item, models.EcomError) {
	var items []entities.Item
	items, err := p.Store.GetAllProducts()
	if err.Message != nil {
		return items, err
	}
	return items, models.EcomError{}
}
