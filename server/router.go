package server

import (
	"github.com/akkinasrikar/ecommerce-cart/controllers"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	servicesLogin "github.com/akkinasrikar/ecommerce-cart/services/login"
	validatorsLogin "github.com/akkinasrikar/ecommerce-cart/validators/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setUpRoutes(router *gin.Engine, db *gorm.DB) {
	servicesLogin := servicesLogin.NewLoginService(repositories.RepositoryInterface(repositories.NewRepository(db)))
	LoginHandler := controllers.NewLoginHandler(servicesLogin, validatorsLogin.NewValidator(), db)
	loginHandler(router, *LoginHandler)
}

func loginHandler(router *gin.Engine, LoginHandler controllers.LoginHandler) {
	router.POST("/signup", LoginHandler.SignUp)
}
