package services

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
)

func (s *loginService) SignUp(req entities.SignUp) (responses.SingUp, error) {
	userDetails, err := s.repoService.SignUp(req)
	if err != nil {
		return responses.SingUp{}, err
	}
	return responses.SingUp{
		Name: userDetails.Name,
		Email: userDetails.Email,
		Message: "User created successfully",
	}, nil
}
