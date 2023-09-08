package services

import (
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/akkinasrikar/ecommerce-cart/validators/helper"
	redis "github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *loginService) SignUp(req models.SignUp) (responses.SingUp, models.EcomError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.SingUp{}, *helper.ErrorInternalSystemError("Error while hashing password : " + err.Error())
	}
	userReq := entities.SignUp{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		UserId:   int64(utils.GenerateRandomUserIdNumber()),
	}
	userDetails, ecomErr := s.repoService.SignUp(userReq)
	if ecomErr.Message != nil {
		return responses.SingUp{}, *helper.ErrorInternalSystemError("Error while signing up : " + ecomErr.Message.Error())
	}
	return responses.SingUp{
		Name:    userDetails.Name,
		Email:   userDetails.Email,
		Message: "User created successfully",
	}, models.EcomError{}
}

func (s *loginService) Login(req models.Login) (responses.Login, models.EcomError) {
	var loginDetails responses.Login
	loginDetails.UserName = req.Name
	userReq := entities.Login{
		Name:     req.Name,
		Password: req.Password,
	}
	userDetails, ecomErr := s.repoService.Login(userReq)
	if ecomErr.Message != nil {
		return responses.Login{}, *helper.ErrorInternalSystemError("Error while logging in : " + ecomErr.Message.Error())
	}
	token, err := s.redisClient.Get(req.Name).Result()
	if err == redis.Nil {
		return responses.Login{}, *helper.ErrorInternalSystemError("Error while getting token from redis cache : " + err.Error())
	}
	if token != "" {
		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("testing"), nil
		})
		if err == nil {
			loginDetails.Token = token
			return loginDetails, models.EcomError{}
		}
	}
	token, err = utils.GenerateToken(req.Name, userDetails.UserId)
	if err != nil {
		return loginDetails, *helper.ErrorInternalSystemError("Error while generating token : " + err.Error())
	}
	loginDetails.Token = token
	err = s.redisClient.Set(req.Name, token, 0).Err()
	if err != nil {
		return loginDetails, *helper.ErrorInternalSystemError("Error while setting token in redis cache : " + err.Error())
	}
	return loginDetails, models.EcomError{}
}
