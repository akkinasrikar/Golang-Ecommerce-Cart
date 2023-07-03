package services

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *loginService) SignUp(req entities.SignUp) (responses.SingUp, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.SingUp{}, err
	}
	req.Password = string(hashedPassword)
	userDetails, err := s.repoService.SignUp(req)
	if err != nil {
		return responses.SingUp{}, err
	}
	return responses.SingUp{
		Name:    userDetails.Name,
		Email:   userDetails.Email,
		Message: "User created successfully",
	}, nil
}

func (s *loginService) Login(req entities.Login) (responses.Login, error) {
	_, err := s.repoService.Login(req)
	if err != nil {
		return responses.Login{}, err
	}
	token, err := utils.GenerateToken(req.Name)
	if err != nil {
		return responses.Login{}, err
	}
	return responses.Login{
		UserName: req.Name,
		Token:    token,
	}, nil
}
