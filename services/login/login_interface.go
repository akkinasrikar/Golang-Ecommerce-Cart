package services

import (
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/models/responses"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
)

type loginService struct {
	repoService repositories.RepositoryInterface
}

func NewLoginService(respoService repositories.RepositoryInterface) LoginService {
	return &loginService{
		repoService: respoService,
	}
}

type LoginService interface {
	SignUp(req entities.SignUp) (responses.SingUp, error)
	Login(req entities.Login) (responses.Login, error)
}
