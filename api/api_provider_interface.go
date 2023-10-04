package api

import (
	"context"
	"net/http"
	"time"

	"github.com/akkinasrikar/ecommerce-cart/api/dto"
	"github.com/akkinasrikar/ecommerce-cart/config"
	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models"
)

//go:generate mockgen -package mocks -source=api_provider_interface.go -destination=mocks/api_provider_mocks.go
type Service interface {
	GetItems(ecomCtx context.Context) (dto.ItemsResponse, models.EcomError)
	SendMail(req models.SendEmailRequest) error
}

type service struct {
	BaseURL string
	http    HttpCall
}

func NewService() Service {
	timeout := constants.Timeout.API_PROVIDERS
	http := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	return &service{
		BaseURL: config.FakeStore.BaseUrl,
		http:    http,
	}
}

//go:generate mockgen -package mocks -source=api_provider_interface.go -destination=mocks/api_provider_mocks.go
type HttpCall interface {
	Do(req *http.Request) (*http.Response, error)
}
