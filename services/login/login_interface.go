package services

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	"github.com/go-redis/redis"
)

type loginService struct {
	repoService repositories.RepositoryInterface
	redisClient *redis.Client
}

func NewLoginService(respoService repositories.RepositoryInterface, redisClient *redis.Client) LoginService {
	return &loginService{
		repoService: respoService,
		redisClient: redisClient,
	}
}

type LoginService interface {
	SignUp(req entities.SignUp) (responses.SingUp, error)
	Login(req entities.Login) (responses.Login, error)
}
