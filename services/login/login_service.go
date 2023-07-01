package services

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
)

func (s *loginService) SignUp(req entities.SignUp) (responses.SingUp, error) {
	return responses.SingUp{
		Name:  req.Name,
		Email: req.Email,
		Message: "Successfully signed up",
	}, nil

}
