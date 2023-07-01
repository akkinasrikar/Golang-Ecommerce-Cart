package services

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
)

type loginService struct {}


func NewLoginService() LoginService {
	return &loginService{}
}

type LoginService interface {
	SignUp(req entities.SignUp) (responses.SingUp, error)
}

