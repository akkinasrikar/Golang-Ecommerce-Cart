package services

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	redis "github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
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
	var loginDetails responses.Login
	loginDetails.UserName = req.Name
	_, err := s.repoService.Login(req)
	if err != nil {
		return responses.Login{}, err
	}
	token, err := s.redisClient.Get(req.Name).Result()
	if err == redis.Nil {
		return responses.Login{}, err
	}
	if token != "" {
		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("testing"), nil
		})
		if err == nil {
			loginDetails.Token = token
			return loginDetails, nil
		}
	}
	token, err = utils.GenerateToken(req.Name)
	if err != nil {
		return loginDetails, err
	}
	loginDetails.Token = token
	err = s.redisClient.Set(req.Name, token, 0).Err()
	if err != nil {
		return loginDetails, err
	}
	return loginDetails, nil
}
